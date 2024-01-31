package config

import (
	"fmt"
	"strings"
	"unicode"
)

type TagPartNotRecognizedError struct {
	Part string
}

func (err *TagPartNotRecognizedError) Error() string {
	return fmt.Sprintf("Tag part %s not recognized", err.Part)
}

type tag struct {
	Name              string
	Default           string
	Required          bool
	NoStructRecursive bool
}

func parseTag(s string) tag {
	t := tag{}

	parts := strings.Split(s, ",")

	t.Name = parts[0]

	for _, part := range parts[1:] {
		switch part {
		case "required", "require", "req":
			t.Required = true
		case "nr", "norecursive", "no-recursive", "no-rec", "nonrecursive", "non-recursive", "non-rec":
			t.NoStructRecursive = true
		default:
			if strings.HasPrefix(part, "default=") {
				t.Default = strings.TrimPrefix(part, "default=")
			} else {
				panic(&TagPartNotRecognizedError{
					Part: part,
				})
			}
		}
	}

	return t
}

// formatName formats the name of a struct field to be used as an env variable name
// examples:
// - "DBName" -> "DB_NAME"
// - "DB" -> "DB"
// - "Name" -> "NAME"
// - "AllowedOrigins" -> "ALLOWED_ORIGINS"
// - "TestA" -> "TEST_A"
func formatName(name string) string {
	var result strings.Builder
	l := len(name) - 1
	for i, r := range name {
		// if current letter is uppercase, not first,
		//   and next letter is lowercase
		//   or if letter is last and previous letter is lowercase:
		//   insert _ before current letter
		if unicode.IsUpper(r) && i != 0 &&
			(i < l && unicode.IsLower(rune(name[i+1])) ||
				i == l && unicode.IsLower(rune(name[i-1]))) {
			result.WriteRune('_')
		}
		result.WriteRune(r)
	}
	return strings.ToUpper(result.String())
}
