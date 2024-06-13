package analyzer

import (
	"golang.org/x/net/html"
	"net/url"
	"strings"
)

// extractLinks traverses the HTML text and sort the found links according to their character: internal or external
// it currently looks for links in <a> and <link> tags, but could be extended to include more tags, such as <script> or <img>
func extractLinks(body, domain string) ([]string, []string) {
	var baseURL *url.URL
	var intLinks, extLinks []string

	tknzr := html.NewTokenizer(strings.NewReader(body))
	for {
		tokType := tknzr.Next()

		if tokType == html.ErrorToken {
			break // stop looping at the end of the text
		}

		if tokType == html.StartTagToken ||
			tokType == html.SelfClosingTagToken { // the base and the link tag can be self-closing

			tok := tknzr.Token()
			switch tok.Data {
			case baseTag: // N.B. this code assumes that - if present - there may be only one base tag in the HTML text
				for _, attr := range tok.Attr {
					link := attr.Val
					if attr.Key == hrefAttr {
						parsedBaseURL, err := url.Parse(link) // discarding invalid URLs
						if err != nil {
							continue
						}
						baseURL = parsedBaseURL
					}
				}
			case aTag, linkTag:
				for _, attr := range tok.Attr {
					if attr.Key == hrefAttr {
						link := attr.Val
						parsedLink, err := url.Parse(link) // discarding formally invalid URLs
						if err != nil {
							continue
						}
						if parsedLink.Scheme == mailtoScheme || // discarding mail addresses, telephone numbers and js scripts
							parsedLink.Scheme == telScheme ||
							parsedLink.Scheme == jsScheme {
							continue
						}
						if baseURL != nil && !parsedLink.IsAbs() {
							link = addBaseToLink(baseURL, link)
						}
						if isInternalLink(parsedLink, domain) {
							intLinks = append(intLinks, link)

						} else {
							extLinks = append(extLinks, link)
						}
					}
				}
			}
		}
	}

	return intLinks, extLinks
}
