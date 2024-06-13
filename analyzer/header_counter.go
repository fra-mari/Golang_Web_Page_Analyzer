package analyzer

import (
	"golang.org/x/net/html"
	"regexp"
	"strings"
)

// countHeaders counts the number of headers in the HTML document, sorted by type.
// It accepts the body of the HTML document as a string as an input and returns a map of header types to their respective counts.
// N.B. The function uses a regular expression to match header tags and assumes the input to be UTF-8 encoded.
func countHeaders(body string) map[string]int {
	headers := make(map[string]int)
	re := regexp.MustCompile(headerTagsRE) // it is safe to compile with the provided constant

	tknzr := html.NewTokenizer(strings.NewReader(body))
	for {
		tokType := tknzr.Next()

		if tokType == html.ErrorToken {
			break // stop looping at the end of the text
		}

		if tokType == html.StartTagToken {
			tok := tknzr.Token()
			if re.MatchString(tok.Data) {
				headerType := strings.ToLower(tok.Data)
				headers[headerType]++
			}
		}
	}

	return headers
}
