package config

import (
	"reflect"
	"strconv"
	"strings"
)

func setField(field reflect.StructField, tg tag, v reflect.Value, val string, found bool) error {
	if !v.CanSet() || val == "" {
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

	switch t.Kind() {
	case reflect.String:
		v.SetString(val)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		cV, err := strconv.Atoi(val)
		if err != nil {
			return err
		}
		v.SetInt(int64(cV))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		cV, err := strconv.Atoi(val)
		if err != nil {
			return err
		}
		v.SetUint(uint64(cV))
	case reflect.Bool:
		cV, err := strconv.ParseBool(val)
		if err != nil {
			return err
		}
		v.SetBool(cV)
	case reflect.Float32:
		cV, err := strconv.ParseFloat(val, 32)
		if err != nil {
			return err
		}
		v.SetFloat(cV)
	case reflect.Float64:
		cV, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return err
		}
		v.SetFloat(cV)
	case reflect.Slice:
		switch field.Type.Elem().Kind() {
		case reflect.String:
			v.Set(reflect.ValueOf(strings.Split(val, " ")))
		case reflect.Int:
			var ints []int
			for _, s := range strings.Split(val, " ") {
				cV, err := strconv.Atoi(s)
				if err != nil {
					return err
				}
				ints = append(ints, cV)
			}
			v.Set(reflect.ValueOf(ints))
		case reflect.Bool:
			var bools []bool
			for _, s := range strings.Split(val, " ") {
				cV, err := strconv.ParseBool(s)
				if err != nil {
					return err
				}
				bools = append(bools, cV)
			}
			v.Set(reflect.ValueOf(bools))
		case reflect.Float32:
			var floats []float32
			for _, s := range strings.Split(val, " ") {
				cV, err := strconv.ParseFloat(s, 32)
				if err != nil {
					return err
				}
				floats = append(floats, float32(cV))
			}
			v.Set(reflect.ValueOf(floats))
		case reflect.Float64:
			var floats []float64
			for _, s := range strings.Split(val, " ") {
				cV, err := strconv.ParseFloat(s, 64)
				if err != nil {
					return err
				}
				floats = append(floats, cV)
			}
			v.Set(reflect.ValueOf(floats))
		default:
			return UnsupportedTypeError{
				Type:  field.Type.Elem().Kind().String(),
				Field: tg.Name,
			}
		}
	default:
		return UnsupportedTypeError{
			Type:  field.Type.Kind().String(),
			Field: tg.Name,
		}
	}

	return nil
}
