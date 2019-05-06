package sitemap

import (
	"fmt"
	"io"
)

const (
	sitemapHeader = `<?xml version="1.0" encoding="UTF-8"?><urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`
	sitemapFooter = `</urlset>`
	urlTpl        = `<url><loc>%s</loc></url>`
)

// Build sitemap by urls and write it to writer.
func Build(urls []string, writer io.Writer) {
	if len(urls) == 0 {
		return
	}

	writer.Write([]byte(sitemapHeader))
	for _, url := range urls {
		part := fmt.Sprintf(urlTpl, url)
		writer.Write([]byte(part))
	}

	writer.Write([]byte(sitemapFooter))
}
