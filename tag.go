package config

import "strings"

type tag struct {
	Name     string
	Default  string
	Required bool
}

func parseTag(s string) tag {
	t := tag{}

	parts := strings.Split(s, ",")

	t.Name = parts[0]

	for _, part := range parts[1:] {
		switch part {
		case "required":
			t.Required = true
		default:
			if strings.HasPrefix(part, "default=") {
				t.Default = strings.TrimPrefix(part, "default=")
			}
		}
	}

	return t
}
