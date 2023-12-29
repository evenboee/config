package config

import (
	"reflect"
	"strconv"
	"strings"
)

func setField(field reflect.StructField, tg tag, v reflect.Value, val string) error {
	switch field.Type.Kind() {
	case reflect.String:
		v.SetString(val)
	case reflect.Int:
		cV, err := strconv.Atoi(val)
		if err != nil {
			return err
		}
		v.SetInt(int64(cV))
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
