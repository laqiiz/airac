package model

import "errors"

var NotFound = errors.New("not found")

// Problem Details for HTTP APIs
// https://tools.ietf.org/html/rfc7807
type ProblemError struct {
	Type     string   `json:"type,omitempty"`
	Title    string   `json:"title"`
	Detail   string   `json:"detail,omitempty"`
	Instance string   `json:"instance,omitempty"`
}
