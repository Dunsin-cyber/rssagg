package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API Key from
// the header of an HTTP request.
// Exmaple
//Authorization: ApiKey {insert api key here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authorization header found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2  {
		return "", errors.New("malformed authorization header ")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("invalid authorization header prefix")
	}
	return vals[1], nil
}