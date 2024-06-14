package models

import "html/template"

type AnalysisResult struct {
	HTMLVersion       string
	PageTitle         string
	HeadersCount      map[string]int
	InternalLinks     int
	ExternalLinks     int
	InaccessibleLinks int
	HasLoginForm      bool
	ErrorMessage      template.HTML // safe to use as all the HTML comes inside the app itself
}
