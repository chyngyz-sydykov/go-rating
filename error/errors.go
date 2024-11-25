package my_error

import "errors"

var ErrNotFound = errors.New("resource not found")
var ErrInvalidArgument = errors.New("resource is not valid")
