package error

import "errors"

var (
	ErrNotFound = errors.New("Not Found")
	ErrInvalid  = errors.New("Invalid")
)
