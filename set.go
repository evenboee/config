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
		v.SetInt(int64(Must(strconv.Atoi(val))))
	case reflect.Bool:
		v.SetBool(Must(strconv.ParseBool(val)))
	case reflect.Float32:
		v.SetFloat(Must(strconv.ParseFloat(val, 32)))
	case reflect.Float64:
		v.SetFloat(Must(strconv.ParseFloat(val, 64)))
	case reflect.Slice:
		switch field.Type.Elem().Kind() {
		case reflect.String:
			v.Set(reflect.ValueOf(strings.Split(val, " ")))
		case reflect.Int:
			var ints []int
			for _, s := range strings.Split(val, " ") {
				ints = append(ints, Must(strconv.Atoi(s)))
			}
			v.Set(reflect.ValueOf(ints))
		case reflect.Bool:
			var bools []bool
			for _, s := range strings.Split(val, " ") {
				bools = append(bools, Must(strconv.ParseBool(s)))
			}
			v.Set(reflect.ValueOf(bools))
		case reflect.Float32:
			var floats []float32
			for _, s := range strings.Split(val, " ") {
				floats = append(floats, float32(Must(strconv.ParseFloat(s, 32))))
			}
			v.Set(reflect.ValueOf(floats))
		case reflect.Float64:
			var floats []float64
			for _, s := range strings.Split(val, " ") {
				floats = append(floats, Must(strconv.ParseFloat(s, 64)))
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
