package analyzer

import (
	"net/url"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

// HTMLAnalyzer defines an interface for analyzing HTML content.
type (
	HTMLAnalyzer interface {
		getHTMLVersion(string) string
		countHeaders(string) map[string]int
		extractTitle(string) string
		detectLoginForm(string) bool
		extractLinks(string, string) ([]string, []string)
	}

	// htmlAnalyzer is a struct that implements the HTMLAnalyzer interface
	htmlAnalyzer struct{}
)

var htmlDeclarations = map[string]string{
	"HTML 1.0":               `"-//IETF//DTD HTML 1.0//EN"`,
	"HTML 2.0":               `"-//IETF//DTD HTML 2.0//EN"`,
	"HTML 3.2":               `"-//W3C//DTD HTML 3.2//EN"`,
	"HTML 4.0 Strict":        `"-//W3C//DTD HTML 4.0//EN"`,
	"HTML 4.0 Transitional":  `"-//W3C//DTD HTML 4.0 Transitional//EN"`,
	"HTML 4.0 Frameset":      `"-//W3C//DTD HTML 4.0 Frameset//EN"`,
	"HTML 4.01 Strict":       `"-//W3C//DTD HTML 4.01//EN"`,
	"HTML 4.01 Transitional": `"-//W3C//DTD HTML 4.01 Transitional//EN"`,
	"HTML 4.01 Frameset":     `"-//W3C//DTD HTML 4.01 Frameset//EN"`,
	"XHTML 1.0 Strict":       `"-//W3C//DTD XHTML 1.0 Strict//EN"`,
	"XHTML 1.0 Transitional": `"-//W3C//DTD XHTML 1.0 Transitional//EN"`,
	"XHTML 1.0 Frameset":     `"-//W3C//DTD XHTML 1.0 Frameset//EN"`,
	"XHTML 1.1":              `"-//W3C//DTD XHTML 1.1//EN"`,
	"HTML 5":                 `<!DOCTYPE html>`,
}

// countHeaders counts the number of headers in the HTML document, sorted by type.
// It accepts the body of the HTML document as a string as an input and returns a map of header types to their respective counts.
// N.B. The function uses a regular expression to match header tags and assumes the input to be UTF-8 encoded.
func (ha htmlAnalyzer) countHeaders(body string) map[string]int {
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

// getHTMLVersion returns the version of the HTML document as a string. It accepts the body of the HTML document as input.
func (ha htmlAnalyzer) getHTMLVersion(body string) string {
	version := unknownVersion

	for name, declaration := range htmlDeclarations {
		if strings.Contains(body, strings.ToLower(declaration)) ||
			strings.Contains(body, strings.ToUpper(declaration)) ||
			strings.Contains(body, declaration) {
			version = name
			break
		}
	}

	return version
}

// extractLinks traverses the HTML text and sort the found links according to their character: internal or external
// it currently looks for links in <a> and <link> tags, but could be extended to include more tags, such as <script> or <img>
func (ha htmlAnalyzer) extractLinks(body, domain string) ([]string, []string) {
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

// detectLoginForm tries to detect the presence of a login form in the HTML document.
// It accepts the body of the HTML document as a string as input and returns a boolean value.
// It looks both for the presence of a password input field and a form with an action attribute set to "login".
func (ha htmlAnalyzer) detectLoginForm(body string) bool {
	var hasLoginForm bool
	tknzr := html.NewTokenizer(strings.NewReader(body))
	for {
		tokType := tknzr.Next()

		if tokType == html.ErrorToken {
			break // stop looping at the end of the text
		}

		if tokType == html.StartTagToken ||
			tokType == html.SelfClosingTagToken { // the input tag can be self-closing
			tok := tknzr.Token()
			switch tok.Data {
			case inputTag:
				if hasAttribute(tok, typeAttr) == passwordAttrVal {
					hasLoginForm = true
					break
				}
			case formTag:
				if hasAttribute(tok, actionAttr) == loginAttrVal {
					hasLoginForm = true
					break
				}
				// Check for CSS classes as well
				if hasAttribute(tok, classAttr) == loginFormCSSAttrVal {
					hasLoginForm = true
					break
				}
			}
		}
	}

	return hasLoginForm
}

// extractTitle extracts the title of the HTML document.
// It accepts the body of the HTML document as a string as input and returns the title as a string.
// N.B. The function assumes the input to be UTF-8 encoded and implies that the page has a single title tag or that the first one is the most relevant.
func (ha htmlAnalyzer) extractTitle(body string) string {
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

func newHTMLAnalyzer() HTMLAnalyzer {
	return htmlAnalyzer{}
}
