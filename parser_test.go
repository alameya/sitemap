package sitemap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	var expectedUrls = []string{
		"faq.php",
		"protocol.php",
		"protocol.php",
		"http://creativecommons.org/licenses/by-sa/2.5/",
		"terms.php",
	}

	actualUrls, err := Parse([]byte(testPage))

	assert.NoError(t, err)
	assert.Equal(t, len(expectedUrls), len(actualUrls))

	for _, url := range expectedUrls {
		assert.Contains(t, actualUrls, url)
	}
}

func TestParseWithBaseTag(t *testing.T) {
	var expectedUrls = []string{
		"http://example.com/testpage.php",
		"http://externallink.com/testpage.php",
		"http://example.com/faq.php",
		"http://example.com/doubleslash",
		"http://example.com/protocol.php",
	}

	actualUrls, err := Parse([]byte(testPageWithBaseTag))

	assert.NoError(t, err)
	assert.Equal(t, len(expectedUrls), len(actualUrls))

	for _, url := range expectedUrls {
		assert.Contains(t, actualUrls, url)
	}
}
