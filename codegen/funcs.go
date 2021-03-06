package codegen

import "strings"

// pascalCase converts valid SNAKE_CASE string to pascalCase string.
func pascalCase(s string) string {
	// TODO: handle invalid input.
	parts := strings.Split(s, "_")
	s = ""
	for _, part := range parts {
		s += strings.Title(strings.ToLower(part))
	}
	return s
}

func replace(s string) string {
	return strings.Replace(s, "\n", "\\n", -1)
}
