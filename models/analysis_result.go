package models

type AnalysisResult struct {
	HTMLVersion       string
	PageTitle         string
	HeadersCount      map[string]int
	InternalLinks     int
	ExternalLinks     int
	InaccessibleLinks int
	HasLoginForm      bool
	ErrorMessage      string
}
