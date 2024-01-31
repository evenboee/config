package config

import "testing"

func Test__formatName(t *testing.T) {

	// - "DBName" -> "DB_NAME"
	// - "DB" -> "DB"
	// - "Name" -> "NAME"
	// - "AllowedOrigins" -> "ALLOWED_ORIGINS"
	// - "TestA" -> "TEST_A"

	tc := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"DBName", "DB_NAME"},
		{"DB", "DB"},
		{"Name", "NAME"},
		{"AllowedOrigins", "ALLOWED_ORIGINS"},
		{"TestA", "TEST_A"},
		{"TString", "T_STRING"},
		{"TInt8", "T_INT8"},
	}

	for _, c := range tc {
		t.Run(c.input, func(t *testing.T) {
			res := formatName(c.input)
			requireEquals(t, res, c.expected)
		})
	}
}
