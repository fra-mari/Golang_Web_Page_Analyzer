package main

const (
	unknownVersion      = "unknown HTML version"
	headerTagsRE        = "[hH][1-6]"
	titleTag            = "title"
	inputTag            = "input"
	formTag             = "form"
	baseTag             = "base"
	aTag                = "a"
	linkTag             = "link"
	typeAttr            = "type"
	actionAttr          = "action"
	classAttr           = "class"
	hrefAttr            = "href"
	loginAttrVal        = "login"
	loginFormCSSAttrVal = "login-form"
	passwordAttrVal     = "password"
	mailtoScheme        = "mailto"
	telScheme           = "tel"
	jsScheme            = "javascript"
	httpStr             = "http"
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
