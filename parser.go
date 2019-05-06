package sitemap

import (
	"regexp"
	"strings"
)

var (
	linkRe  = regexp.MustCompile(`<a.+?href="([^"]+)`)
	baseRe  = regexp.MustCompile(`<base.+?href="([^"]+)`)
	ancorRe = regexp.MustCompile(`#[^?"]*`)
)

// Parse extract urls from body.
// If <base href="..."> tag is present add href value before url.
func Parse(body []byte) ([]string, error) {
	var baseURL string
	baseTagResult := baseRe.FindAllSubmatch(body, 1)
	if len(baseTagResult) != 0 && len(baseTagResult[0]) == 2 {
		baseURL = string(baseTagResult[0][1])
	}

	matches := linkRe.FindAllSubmatch(body, -1)
	var urls []string

	for _, match := range matches {
		url := string(ancorRe.ReplaceAll(match[1], []byte("")))
		if url == "" {
			continue
		}

		switch {
		case strings.HasPrefix(url, "//"): // for uniqueUrls which starts from "//" // TODO check
			url = "http:" + url
		case !strings.HasPrefix(url, "http"): // relative paths
			if baseURL != "" {
				if strings.HasPrefix(url, "/") {
					url = baseURL + url
				} else {
					url = baseURL + "/" + url
				}
			}                                 // else TODO implement autocomplete relative uniqueUrls
		}

		// TODO ignore external domains

		urls = append(urls, url)
	}

	return urls, nil
}
