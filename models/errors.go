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
		return fmt.Sprintf("HTTP status code: %d - %s", e.Code, m)
	}

	return fmt.Sprintf("HTTP status code: %d", e.Code)
}

var errorMessages = map[int]string{
	http.StatusMovedPermanently:    "Moved Permanently: The content you are looking for has been moved to a new location.",
	http.StatusFound:               "Found: The content has been found, but it temporarily resides at a different location.",
	http.StatusBadRequest:          "Bad Request: The request was not formed correctly.",
	http.StatusUnauthorized:        "Unauthorized: Authentication credentials are required to access this content.",
	http.StatusForbidden:           "Forbidden: Access to the requested resource is denied.",
	http.StatusNotFound:            "Not Found: The requested content could not be found.",
	http.StatusTooManyRequests:     "Too Many Requests. Please try again later.",
	http.StatusInternalServerError: "Internal Server Error: An error has occurred on the server side.",
	http.StatusBadGateway:          "Bad Gateway: Received an invalid response from the upstream server.",
	http.StatusServiceUnavailable:  "Service Unavailable: The server is currently unable to handle the request.",
	http.StatusGatewayTimeout:      "Gateway Timeout: The server acted as a gateway or proxy and did not receive a timely response from the upstream server.",
}
