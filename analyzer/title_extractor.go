package analyzer

import (
	"golang.org/x/net/html"
	"strings"
)

// extractTitle extracts the title of the HTML document.
// It accepts the body of the HTML document as a string as input and returns the title as a string.
// N.B. The function assumes the input to be UTF-8 encoded and implies that the page has a single title tag or that the first one is the most relevant.
func extractTitle(body string) string {
	var title string
	tknzr := html.NewTokenizer(strings.NewReader(body))
	for {
		tokType := tknzr.Next()

		if tokType == html.ErrorToken {
			break // stop looping at the end of the text
		}

		if tokType == html.StartTagToken {
			tok := tknzr.Token()
			if tok.Data == titleTag {
				tokType = tknzr.Next()
				if tokType == html.TextToken {
					title = tknzr.Token().Data
					break
				}
			}
		}
	}

	return title
}
