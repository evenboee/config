package config

import (
	"reflect"
	"testing"
)

func requireEquals(t *testing.T, a any, b any) {
	if !reflect.DeepEqual(a, b) {
		t.Fatalf("expected:\n\t%+v\nto equal:\n\t%+v\n", a, b)
	}
}

func requireNoErr(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("expected no error, got: (%T) %v", err, err)
	}
}
