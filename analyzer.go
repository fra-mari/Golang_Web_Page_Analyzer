package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"regexp"
	"strings"
)

// getHTMLVersion returns the version of the HTML document as a string. It accepts the body of the HTML document as input.
func getHTMLVersion(body string) string {
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

// countHeaders counts the number of headers in the HTML document, sorted by type.
// It accepts the body of the HTML document as a string as an input and returns a map of header types to their respective counts.
// N.B. The function uses a regular expression to match header tags and assumes the input to be UTF-8 encoded.
func countHeaders(body string) map[string]int {
	headers := make(map[string]int)
	re := regexp.MustCompile(headerTags) // it is safe to compile with the provided constant

	tok := html.NewTokenizer(strings.NewReader(body))
	for {
		tokType := tok.Next()
		if tokType == html.ErrorToken {
			break // stop looping at the end of the text
		}
		if tokType == html.StartTagToken {
			tag := tok.Token()
			if re.MatchString(tag.Data) {
				headerType := strings.ToLower(tag.Data)
				headers[headerType]++
			}
		}
	}

	return headers
}

// extractTitle extracts the title of the HTML document. It accepts the body of the HTML document as a string as input and returns the title as a string.
// N.B. The function assumes the input to be UTF-8 encoded and implies that the page has a single title tag or that the first one is the most relevant.
func extractTitle(body string) string {
	var title string
	tok := html.NewTokenizer(strings.NewReader(body))
	for {
		tokType := tok.Next()
		if tokType == html.ErrorToken {
			break // stop looping at the end of the text
		}
		if tokType == html.StartTagToken {
			tag := tok.Token()
			if tag.Data == titleTag {
				tokType = tok.Next()
				if tokType == html.TextToken {
					title = tok.Token().Data
					break
				}
			}
		}
	}

	return title
}

// detectLoginForm tries to detect the presence of a login form in the HTML document.
// It accepts the body of the HTML document as a string as input and returns a boolean value.
// It looks both for the presence of a password input field and a form with an action attribute set to "login".
func detectLoginForm(body string) bool {
	var hasLoginForm bool
	tok := html.NewTokenizer(strings.NewReader(body))
	for {
		tokType := tok.Next()
		if tokType == html.ErrorToken {
			break // stop looping at the end of the text
		}
		if tokType == html.StartTagToken ||
			tokType == html.SelfClosingTagToken { // input tag can be self-closing
			tag := tok.Token()
			switch tag.Data {
			case inputTag:
				if hasAttribute(tag, typeAttr) == passwordAttrVal {
					hasLoginForm = true
					break
				}
			case formTag:
				if hasAttribute(tag, actionAttr) == loginAttrVal {
					hasLoginForm = true
					break
				}
				// Check for CSS classes as well
				if hasAttribute(tag, classAttr) == loginFormCSSAttrVal {
					hasLoginForm = true
					break
				}
			}
		}
	}

	return hasLoginForm
}

func AnalyzePage(pageURL string) (AnalysisResult, error) {
	var result AnalysisResult
	resp, err := http.Get(pageURL)
	if err != nil {
		return result, fmt.Errorf("URL is not reachable: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("HTTP status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, fmt.Errorf("error reading HTML: %v", err)
	}
	result.HTMLVersion = getHTMLVersion(string(body))
	result.HeadersCount = countHeaders(string(body))
	result.PageTitle = extractTitle(string(body))
	result.HasLoginForm = detectLoginForm(string(body))

	resp2, err := http.Get(pageURL)
	if err != nil {
		return result, fmt.Errorf("URL is not reachable: %v", err)
	}
	defer resp2.Body.Close()

	doc, err := html.Parse(resp2.Body)
	if err != nil {
		return result, fmt.Errorf("error parsing HTML: %v", err)
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {

			case "a":
				for _, attr := range n.Attr {
					if attr.Key == "href" {
						link := attr.Val
						if strings.HasPrefix(link, "/") || strings.HasPrefix(link, pageURL) {
							result.InternalLinks++
						} else {
							result.ExternalLinks++
						}
						if !isLinkAccessible(link) {
							result.InaccessibleLinks++
						}
						break
					}
				}
			}
		}
		for child := n.FirstChild; child != nil; child = child.NextSibling {
			f(child)
		}
	}

	// Analyze the document
	f(doc)

	return result, nil
}

func isLinkAccessible(link string) bool {
	resp, err := http.Get(link)
	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}
	return true
}
