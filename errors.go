package config

import "errors"

type MissingRequiredFieldError struct {
	VarName    string
	EnvVarName string
	Field      string
}

func (e MissingRequiredFieldError) Error() string {
	d := ""
	if e.EnvVarName != "" {
		d += ", env: " + e.EnvVarName
	}
	if e.VarName != "" {
		d += ", var: " + e.VarName
	}

	return "missing required field: [name: " + e.Field + d + "]"
}

type UnsupportedTypeError struct {
	Type  string
	Field string
}

func (e UnsupportedTypeError) Error() string {
	return "unsupported type: " + e.Type + " for field: " + e.Field
}

// var ErrFieldIsNotString = errors.New("field is not a string")
var ErrTypeIsNotStruct = errors.New("type is not a struct")
