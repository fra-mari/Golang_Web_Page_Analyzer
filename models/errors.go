package models

import (
	"fmt"
	"net/http"
)

type HTTPError struct {
	Code int
}

// Error implements the error interface for HTTPError.
func (e HTTPError) Error() string {
	if m, ok := errorMessages[e.Code]; ok {
		return fmt.Sprintf(`<b>HTTP status code: </b><span class="error-code">%d – </span>%s`, e.Code, m)
	}

	return fmt.Sprintf(`<b>HTTP status code: </b><span class="error-code">%d</span>`, e.Code)
}

var errorMessages = map[int]string{
	http.StatusMovedPermanently:    `<span class="error-code">Moved Permanently</span><b>:</b> The content you are looking for has been moved to a new location.`,
	http.StatusFound:               `<span class="error-code">Found</span><b>:</b> The content has been found, but it temporarily resides at a different location.`,
	http.StatusBadRequest:          `<span class="error-code">Bad Request</span><b>:</b> The request was not formed correctly.`,
	http.StatusUnauthorized:        `<span class="error-code">Unauthorized</span><b>:</b> Authentication credentials are required to access this content.`,
	http.StatusForbidden:           `<span class="error-code">Forbidden</span><b>:</b> Access to the requested resource is denied.`,
	http.StatusNotFound:            `<span class="error-code">Not Found</span><b>:</b> The requested content could not be found.`,
	http.StatusTooManyRequests:     `<span class="error-code">Too Many Requests</span>. Please try again later.`,
	http.StatusInternalServerError: `<span class="error-code">Internal Server Error</span><b>:</b> An error has occurred on the server side.`,
	http.StatusBadGateway:          `<span class="error-code">Bad Gateway</span><b>:</b> Received an invalid response from the upstream server.`,
	http.StatusServiceUnavailable:  `<span class="error-code">Service Unavailable</span><b>:</b> The server is currently unable to handle the request.`,
	http.StatusGatewayTimeout:      `<span class="error-code">Gateway Timeout</span><b>:</b> The server acted as a gateway or proxy and did not receive a timely response from the upstream server.`,
}
