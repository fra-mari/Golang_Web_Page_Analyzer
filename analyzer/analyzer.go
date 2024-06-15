package analyzer

import (
	"fmt"
	"home24/models"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
)

type (
	// Analyzer defines an interface for analyzing a webpage.
	Analyzer interface {
		AnalyzePage(string) (models.AnalysisResult, error)
	}
	// analyzer is a struct that implements the Analyzer interface
	analyzer struct {
		linkChecker  LinkChecker
		htmlAnalyzer HTMLAnalyzer
	}
)

func (a analyzer) AnalyzePage(pageURL string) (models.AnalysisResult, error) {
	var result models.AnalysisResult

	// Validate the URL
	_, err := url.ParseRequestURI(pageURL)
	if err != nil {
		return result, fmt.Errorf("The provided URL is <u>invalid</u>. Please insert a valid URL and try again.")
	}

	resp, err := http.Get(pageURL)
	if err != nil {
		return result, fmt.Errorf("The provided URL is <u>not reachable</u>.<br>(%v)", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return result, models.HTTPError{Code: resp.StatusCode}
	}

	domain := resp.Request.Host

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, fmt.Errorf("Error reading HTML.<br>(%v)", err)
	}

	result.HTMLVersion = a.htmlAnalyzer.getHTMLVersion(string(body))
	result.HeadersCount = a.htmlAnalyzer.countHeaders(string(body))
	result.PageTitle = a.htmlAnalyzer.extractTitle(string(body))
	result.HasLoginForm = a.htmlAnalyzer.detectLoginForm(string(body))

	internals, externals := a.htmlAnalyzer.extractLinks(string(body), domain)
	result.InternalLinks = len(internals)
	result.ExternalLinks = len(externals)

	links := append(internals, externals...)

	var counter int32
	cleanLinks := make([]string, 0, len(links))

	// this is an extra loop, but should increase performance a bit by reducing the number of calls
	for _, l := range links {
		cleanLink := strings.TrimSpace(l)
		if !strings.HasPrefix(cleanLink, httpStr) {
			counter++
			continue
		}
		cleanLinks = append(cleanLinks, cleanLink)
	}

	var wg sync.WaitGroup
	wg.Add(len(cleanLinks)) // Number of goroutines to launch for testing the links accessibility in parallel
	results := make(chan bool, len(cleanLinks))
	for _, cl := range cleanLinks {
		go a.linkChecker.isLinkAccessible(cl, &wg, results)
	}

	go func() {
		for r := range results {
			if !r {
				atomic.AddInt32(&counter, 1)
			}

		}
	}()

	// Wait for all goroutines to finish and close the channel
	wg.Wait()
	close(results)

	result.InaccessibleLinks = counter

	return result, nil
}

func NewAnalyzer() Analyzer {
	return analyzer{
		linkChecker:  newLinkChecker(),
		htmlAnalyzer: newHTMLAnalyzer(),
	}
}
