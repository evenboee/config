package config

import "errors"

type MissingRequiredFieldError struct {
	Field string
}

func (e MissingRequiredFieldError) Error() string {
	return "missing required field: " + e.Field
}

var ErrFieldIsNotString = errors.New("field is not a string")
var ErrTypeIsNotStruct = errors.New("type is not a struct")
