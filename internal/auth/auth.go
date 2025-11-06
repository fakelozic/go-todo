package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts and API Key from the headers of an HTTP request
// Authorization: ApiKey {insert apikey here}
func GetAPIKey(headers http.Header) ([]string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return nil, errors.New("no authentication info found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 3 {
		return nil, errors.New("malformed auth header")
	}
	if vals[0] != "ApiKey" {
		return nil, errors.New("malformed first part of auth header")
	}
	return vals, nil
}
