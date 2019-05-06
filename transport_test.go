package sitemap

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createDefaultTransport(ctx context.Context) *Transport {
	errHandler := func(err error) {}
	const maxDepth = 1
	const parallel = 3

	return NewTransport(
		ctx,
		Parse,
		errHandler,
		maxDepth,
		parallel,
	)
}

func TestNewTransport(t *testing.T) {
	transport := createDefaultTransport(context.Background())
	assert.NotNil(t, transport)
	transport.Close()

	ctx, cancel := context.WithCancel(context.Background())
	transport = createDefaultTransport(ctx)
	assert.NotNil(t, transport)
	cancel()
}

func TestTransport_Fetch(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, testPage)
	}

	testServer := httptest.NewServer(http.HandlerFunc(handler))
	defer testServer.Close()

	transport := createDefaultTransport(context.Background())

	urls := transport.Fetch(testServer.URL)
	assert.Equal(t, 5, len(urls))

	urls = transport.Fetch("")
	assert.Equal(t, 0, len(urls))
}

func TestIncorrectBehavior(t *testing.T) {
	const maxDepth = 1
	const parallel = 1
	var hasError bool

	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, testPage)
	}

	testServer := httptest.NewServer(http.HandlerFunc(handler))
	defer testServer.Close()

	errHandler := func(err error) {}

	parser := func(body []byte) ([]string, error) {
		hasError = true

		return []string{}, errors.New("has error")
	}

	transport := NewTransport(
		context.Background(),
		parser,
		errHandler,
		maxDepth,
		parallel,
	)

	transport.Fetch(testServer.URL)

	assert.True(t, hasError)
}