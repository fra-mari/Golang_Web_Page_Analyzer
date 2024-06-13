package analyzer

import (
	"golang.org/x/net/html"
	"strings"
)

// detectLoginForm tries to detect the presence of a login form in the HTML document.
// It accepts the body of the HTML document as a string as input and returns a boolean value.
// It looks both for the presence of a password input field and a form with an action attribute set to "login".
func detectLoginForm(body string) bool {
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
