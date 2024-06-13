package analyzer

import "strings"

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
