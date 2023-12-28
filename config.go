package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"io/fs"
	"os"
	"reflect"
)

func Guard(err error) {
	if err != nil {
		panic(fmt.Errorf("(%T) %w", err, err))
	}
}

func Must[T any](d T, err error) T {
	Guard(err)
	return d
}

func LoadInto[T any](obj *T, opts ...Option) error {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}

	vals, err := godotenv.Read(o.filenames...)
	if err != nil {
		var pathError *fs.PathError
		if !errors.As(err, &pathError) {
			return err
		} else if !o.ignoreMissingFiles {
			return err
		}
	}

	v := reflect.ValueOf(obj)
	for v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return ErrTypeIsNotStruct
	}

	t := v.Type()
	numFields := t.NumField()
	for i := 0; i < numFields; i++ {
		field := t.Field(i)
		if field.Type.Kind() != reflect.String {
			return ErrFieldIsNotString
		}

		// Skip unexported fields. Panics on set if not exported.
		if !field.IsExported() {
			continue
		}

		tagVal := field.Tag.Get(o.tagName)
		if tagVal == "-" {
			continue
		}

		tg := parseTag(tagVal)
		if tg.Name == "" {
			tg.Name = field.Name
		}

		val := ""
		if !o.omitEnvVars {
			val = os.Getenv(o.envPrefix + tg.Name)
		}

		if val == "" {
			val = vals[o.varPrefix+tg.Name]
		}

		if val == "" && !o.omitDefaults {
			val = o.defaultOverrides[tg.Name]
			if val == "" {
				val = tg.Default
			}
		}

		if val == "" && tg.Required {
			return MissingRequiredFieldError{Field: tg.Name}
		}

		v.Field(i).SetString(val)
	}

	return nil
}

func MustLoadInto[T any](obj *T, opts ...Option) {
	Guard(LoadInto(obj, opts...))
}

func LoadFileInto[T any](filename string, obj *T, opts ...Option) error {
	opts = append(opts, WithFilenames(filename))
	return LoadInto(obj, opts...)
}

func MustLoadFileInto[T any](obj *T, filename string, opts ...Option) {
	opts = append(opts, WithFilenames(filename))
	Guard(LoadFileInto(filename, obj, opts...))
}

func Load[T any](opts ...Option) (*T, error) {
	var obj T
	err := LoadInto(&obj, opts...)
	return &obj, err
}

func MustLoad[T any](opts ...Option) *T {
	return Must(Load[T](opts...))
}

func LoadFile[T any](filename string, opts ...Option) (*T, error) {
	opts = append(opts, WithFilenames(filename))
	return Load[T](opts...)
}

func MustLoadFile[T any](filename string, opts ...Option) *T {
	return Must(LoadFile[T](filename, opts...))
}
