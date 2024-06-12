package main

import "golang.org/x/net/html"

// hasAttribute verifies whether a tag has a specific attribute and returns its value.
func hasAttribute(tag html.Token, key string) string {
	for _, attr := range tag.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}

	return ""
}
