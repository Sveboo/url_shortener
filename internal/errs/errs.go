// Package errs contains internal errors
package errs

import "fmt"

var ErrUrlNotFound = fmt.Errorf("no such URL")
var ErrInvalidUrl = fmt.Errorf("invalid url")
var ErrNotAllowedMethod = fmt.Errorf("method not allowed")
