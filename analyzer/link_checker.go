package analyzer

import (
	"net/http"
	"sync"
)

// LinkChecker defines an interface for checking link accessibility.
type (
	LinkChecker interface {
		isLinkAccessible(string, *sync.WaitGroup, chan<- bool)
	}

	// linkChecker is a struct that implements the LinkChecker interface
	linkChecker struct{}
)

// isLinkAccessible verifies whether a link is accessible by making a GET request and checking the status code
// it returns a boolean through a channel
func (lc linkChecker) isLinkAccessible(link string, wg *sync.WaitGroup, results chan<- bool) {
	defer wg.Done()
	resp, err := http.Get(link)
	if err != nil || resp.StatusCode != http.StatusOK {
		results <- false
	}
	results <- true
}

func newLinkChecker() LinkChecker {
	return linkChecker{}
}
