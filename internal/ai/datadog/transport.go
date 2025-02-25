package datadog

import (
	"fmt"
	"net/http"
)

// Transport is a custom http.RoundTripper that supports adding custom headers
// to the request to the Datadog AI proxy API.
type Transport struct {
	// Email is the email address of the user.
	Email string
}

// RoundTrip implements the http.RoundTripper interface.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("X-Datadog-Org-Id", "2")
	req.Header.Set("X-Datadog-Source", "documentor")
	req.Header.Set("X-Datadog-User-Id", t.Email)

	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return resp, nil
}
