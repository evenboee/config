package config

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Unmarshaller interface {
	Unmarshal(string) error
}

var ParseTimeFormat = time.RFC3339

var (
	ErrUnsupportedType = errors.New("unsupported type")
)

func setField(field reflect.StructField, tg tag, v reflect.Value, val string, found bool) error {
	// If the value cannot be set, no need to continue
	if !v.CanSet() { // || val == ""
		return nil
	}

	switch fV := v.Interface().(type) {
	case Unmarshaller:
		if v.IsNil() {
			n := reflect.New(field.Type.Elem())
			v.Set(n)
			fV = n.Interface().(Unmarshaller)
		}

		err := fV.Unmarshal(val)
		if err != nil {
			return err
		}
		return nil
	}

	t := field.Type
	isPointer := false

	// Check if the type is a pointer, and get its element type
	for t.Kind() == reflect.Pointer {
		isPointer = true
		t = t.Elem()
	}

	// If it's a pointer, we need to create a new value to set
	if isPointer && v.IsNil() {
		// In the case where the type is a pointer
		//   and no value has been explicitly set (config or default),
		//   we should maintain the nil value (skip setting it)
		if !found {
			return nil
		}
		v.Set(reflect.New(t))
	}

	// If it's a pointer, get the actual value it should point to
	if isPointer {
		v = v.Elem()
	}

	q := reflect.New(t).Elem().Interface()
	if t.Kind() == reflect.Slice {
		err := setSliceValue(q, v, val)
		if err != nil {
			if err == ErrUnsupportedType {
				return &UnsupportedTypeError{
					Type:  "[]" + field.Type.Elem().Kind().String(),
					Field: tg.Name,
				}
			}
			return err
		}
	} else {
		err := setValue(q, v, val)
		if err != nil {
			if err == ErrUnsupportedType {
				return &UnsupportedTypeError{
					Type:  field.Type.Kind().String(),
					Field: tg.Name,
				}
			}
			return err
		}
	}

	return nil
}

var SliceSeparator = " "

func setSliceValue(t any, v reflect.Value, val string) error {
	vals := strings.Split(val, SliceSeparator)

	res := reflect.MakeSlice(v.Type(), len(vals), len(vals))

	t = reflect.New(v.Type().Elem()).Elem().Interface()
	for i, val := range vals {
		err := setValue(t, res.Index(i), val)
		if err != nil {
			return err
		}
	}

	v.Set(res)
	return nil
}

func setValue(t any, v reflect.Value, val string) error {
	switch t.(type) {
	case string:
		v.SetString(val)
	case int, int8, int16, int32, int64:
		cVal, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(cVal)
	case uint, uint8, uint16, uint32, uint64:
		cVal, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			return err
		}
		v.SetUint(cVal)
	case bool:
		cVal, err := parseBool(val)
		if err != nil {
			return err
		}
		v.SetBool(cVal)
	case float32, float64:
		cVal, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return err
		}
		v.SetFloat(cVal)
	case time.Duration:
		cVal, err := time.ParseDuration(val)
		if err != nil {
			println("\t\terror", err.Error())
			return err
		}
		v.Set(reflect.ValueOf(cVal))
	case time.Time:
		cVal, err := time.Parse(ParseTimeFormat, val)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(cVal))
	default:
		return ErrUnsupportedType
	}

	return nil
}

// copy of strconv.ParseBool but with "yes" and "no"
func parseBool(s string) (bool, error) {
	switch s {
	case "1", "t", "T", "true", "TRUE", "True", "yes", "YES", "Yes":
		return true, nil
	case "0", "f", "F", "false", "FALSE", "False", "no", "NO", "No":
		return false, nil
	}
	return false, &strconv.NumError{
		// syntax of .Error() is weird for this Func
		//   becomes "strconv.config.parseBool" instead of "config.parseBool"
		//   but using strconv.NumError could make error handling easier
		Func: "config.parseBool",
		Num:  string([]byte(s)),
		Err:  strconv.ErrSyntax,
	}
}
