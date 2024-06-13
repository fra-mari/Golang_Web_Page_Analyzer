package analyzer

import (
	"fmt"
	"home24/models"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func AnalyzePage(pageURL string) (models.AnalysisResult, error) {
	var result models.AnalysisResult

	// Validate the URL
	_, err := url.ParseRequestURI(pageURL)
	if err != nil {
		return result, fmt.Errorf("the provided URL is invalid. Please insert a valid URL and try again")
	}

	resp, err := http.Get(pageURL)
	if err != nil {
		return result, fmt.Errorf("the provided URL is not reachable: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return result, models.HTTPError{Code: resp.StatusCode}
	}

	domain := resp.Request.Host

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, fmt.Errorf("error reading HTML: %v", err)
	}

	result.HTMLVersion = getHTMLVersion(string(body))
	result.HeadersCount = countHeaders(string(body))
	result.PageTitle = extractTitle(string(body))
	result.HasLoginForm = detectLoginForm(string(body))

	internals, externals := extractLinks(string(body), domain)
	result.InternalLinks = len(internals)
	result.ExternalLinks = len(externals)

	links := append(internals, externals...)

	counter := 0
	cleanLinks := make([]string, 0, len(links))
	for _, l := range links {
		cleanLink := strings.TrimSpace(l)
		if !strings.HasPrefix(cleanLink, httpStr) { // TODO: come back on that. It reduces the number of calls but is arbitrary
			counter++
			continue
		}
		cleanLinks = append(cleanLinks, cleanLink)
	}
	wg.Add(len(cleanLinks)) // Number of goroutines to launch for testing the links accessibility in parallel
	for _, cl := range cleanLinks {
		go func(link string) {
			defer wg.Done()
			if !isLinkAccessible(link) {
				counter++
			}
		}(cl)
	}
	// Wait for all goroutines to finish
	wg.Wait()

	result.InaccessibleLinks = counter

	return result, nil
}
