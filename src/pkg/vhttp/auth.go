package vhttp

import (
	"net/http"
	"strings"
)

func ExtractAuthValue(prefix string, str string) (string, error) {
	if str == "" {
		return "", EmptyAuthorizationError
	}

	// Extract token
	tokens := strings.Split(str, " ")
	if len(tokens) != 2 {
		return "", MalformedTokenError
	}

	// Check prefix
	if tokens[0] != prefix {
		return "", MalformedTokenError
	}

	return tokens[1], nil
}

func ExtractBearerAuth(r *http.Request) (token string, err error) {
	// Get header
	authHeader := r.Header.Get(AuthorizationHeader)

	// Extract base64 encoded
	bearerToken, err := ExtractAuthValue("Bearer", authHeader)
	if err != nil {
		return "", EmptyAuthorizationError
	}

	return bearerToken, nil
}
