package config

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"reflect"

	"github.com/joho/godotenv"
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

func loadValsInto(o *config, values map[string]string, obj any) error {
	v := reflect.ValueOf(obj)
	for v.Kind() == reflect.Pointer {
		// Initialize a new instance if the pointer is nil
		// Fixes issue with sub-structs not being initialized
		if v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return ErrTypeIsNotStruct
	}

	t := v.Type()
	numFields := t.NumField()
	for i := 0; i < numFields; i++ {
		field := t.Field(i)

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

			if o.autoFormatFieldName {
				tg.Name = formatName(tg.Name)
			}
		}

		// if fieldT.Kind() == reflect.Struct && !tg.NoStructRecursive {
		// 	if err := loadValsInto(o.Copy().With(
		// 		WithEnvPrefix(o.envPrefix+tg.Name),
		// 		WithVarPrefix(o.varPrefix+tg.Name),
		// 	), values, v.Field(i).Addr().Interface()); err != nil {
		// 		return err
		// 	}
		// 	continue
		// }

		val := ""
		if !o.omitEnvVars {
			val = os.Getenv(o.envPrefix + tg.Name)
		}

		found := val != ""

		if val == "" {
			val, found = values[o.varPrefix+tg.Name]
		}

		if val == "" && !o.omitDefaults {
			val, found = o.defaultOverrides[tg.Name]
			if val == "" {
				val = tg.Default
				found = tg.Default != ""
			}
		}

		if val == "" && tg.Required {
			return MissingRequiredFieldError{
				VarName:    o.varPrefix + tg.Name,
				EnvVarName: o.envPrefix + tg.Name,
				Field:      field.Name,
			}
		}

		if err := setField(field, tg, v.Field(i), val, found); err != nil {
			switch err.(type) {
			case *UnsupportedTypeError:
				fieldT := field.Type
				for fieldT.Kind() == reflect.Pointer {
					fieldT = fieldT.Elem()
				}

				if fieldT.Kind() == reflect.Struct && !tg.NoStructRecursive {
					err = loadValsInto(o.Copy().With(
						WithEnvPrefix(o.envPrefix+tg.Name),
						WithVarPrefix(o.varPrefix+tg.Name),
					), values, v.Field(i).Addr().Interface())
					if err != nil {
						return err
					}
				}
			default:
				return err
			}
		}
	}

	return nil
}

func LoadInto(obj any, opts ...Option) error {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}

	var vals map[string]string
	var err error

	if len(o.filenames) > 0 {
		vals, err = godotenv.Read(o.filenames...)
		if err != nil {
			var pathError *fs.PathError
			if !errors.As(err, &pathError) {
				return err
			} else if !o.ignoreMissingFiles {
				return err
			}
		}
	}

	return loadValsInto(o, vals, obj)
}

func MustLoadInto[T any](obj *T, opts ...Option) {
	Guard(LoadInto(obj, opts...))
}

func LoadFileInto(filename string, obj any, opts ...Option) error {
	opts = append(opts, WithFilenames(filename))
	return LoadInto(obj, opts...)
}

func MustLoadFileInto(obj any, filename string, opts ...Option) {
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
