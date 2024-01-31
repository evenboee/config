package config

import (
	"strconv"
	"testing"
	"time"
)

type UnmarshallerType struct {
	N int
}

func (u *UnmarshallerType) Unmarshal(text string) error {
	n, err := strconv.Atoi(text)
	if err != nil {
		return err
	}

	u.N = 2 * n
	return nil
}

var _ Unmarshaller = &UnmarshallerType{}

func Test__loadValsInto(t *testing.T) {
	type Test struct {
		Abc struct {
			Def struct {
				Ghi []int
			}
		}

		TString string `config:",required"`

		TInt   int   `config:",required"`
		TInt8  int8  `config:",required"`
		TInt16 int16 `config:",required"`
		TInt32 int32 `config:",required"`
		TInt64 int64 `config:",required"`

		TUint   uint   `config:",required"`
		TUint8  uint8  `config:",required"`
		TUint16 uint16 `config:",required"`
		TUint32 uint32 `config:",required"`
		TUint64 uint64 `config:",required"`

		TFloat32 float32 `config:",required"`
		TFloat64 float64 `config:",required"`

		TBool bool `config:",required"`

		TDuration time.Duration `config:",required"`
		TTime     time.Time     `config:",required"`

		TUnmarshaler *UnmarshallerType `config:",required"`
	}

	obj := Test{}
	opts := defaultOptions()
	vals := map[string]string{
		"ABC_DEF_GHI":   "1 2 3",
		"T_STRING":      "abc",
		"T_INT":         "1",
		"T_INT8":        "2",
		"T_INT16":       "3",
		"T_INT32":       "4",
		"T_INT64":       "5",
		"T_UINT":        "6",
		"T_UINT8":       "7",
		"T_UINT16":      "8",
		"T_UINT32":      "9",
		"T_UINT64":      "10",
		"T_FLOAT32":     "11.1",
		"T_FLOAT64":     "12.2",
		"T_BOOL":        "true",
		"T_DURATION":    "1h",
		"T_TIME":        "2006-01-02T15:04:05Z",
		"T_UNMARSHALER": "123",
	}
	expected := Test{
		Abc: struct {
			Def struct {
				Ghi []int
			}
		}{
			Def: struct {
				Ghi []int
			}{
				Ghi: []int{1, 2, 3},
			},
		},
		TString: "abc",
		TInt:    1, TInt8: 2, TInt16: 3, TInt32: 4, TInt64: 5,
		TUint: 6, TUint8: 7, TUint16: 8, TUint32: 9, TUint64: 10,
		TFloat32: 11.1, TFloat64: 12.2,
		TBool:     true,
		TDuration: time.Hour,
		TTime:     time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		TUnmarshaler: &UnmarshallerType{
			N: 246,
		},
	}

	err := loadValsInto(opts, vals, &obj)
	requireNoErr(t, err)

	requireEquals(t, obj, expected)
}
