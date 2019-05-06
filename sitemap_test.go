package sitemap

import (
	"bufio"
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

var testUrls = []string{"https://www.sitemaps.org/", "https://www.sitemaps.org/protocol.html"}

func TestBuild(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)

	Build(testUrls, w)

	w.Flush()

	expected := sitemapHeader +
		"<url><loc>https://www.sitemaps.org/</loc></url>" +
		"<url><loc>https://www.sitemaps.org/protocol.html</loc></url>" +
		sitemapFooter

	got, err := ioutil.ReadAll(&b)
	assert.NoError(t, err)
	assert.Equal(t, expected, string(got))
}

func TestBuildWithoutData(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)

	Build([]string{}, w)

	w.Flush()

	got, err := ioutil.ReadAll(&b)
	assert.NoError(t, err)
	assert.Equal(t, "", string(got))
}
