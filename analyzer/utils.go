package analyzer

import (
	"golang.org/x/net/html"
	"log"
	"net/url"
	"strings"
)

// hasAttribute verifies whether a tag has a specific attribute and returns its value.
func hasAttribute(tag html.Token, key string) string {
	for _, attr := range tag.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}

	return ""
}

// addBaseToLink joins produces an absolute path joining a base with the remaining URL
func addBaseToLink(base *url.URL, link string) string {
	joined, err := base.Parse(link)
	if err != nil {
		log.Printf("error joining base URL %s with link URL %s: %v", base.String(), link, err)
		return link
	}
	return joined.String()
}

// isInternalLink seeks to verify whether the link is internal or not, returning a boolean
func isInternalLink(link *url.URL, domain string) bool {
	if link.Host == "" {
		// if the url's host is not provided, the link is internal
		return true
	}
	if strings.EqualFold(link.Host, domain) {
		// this is an absolute URL but within the same domain, hence corresponds to an internal link
		// N.B. subdomains are treated like external links
		return true
	}

	return false
}
