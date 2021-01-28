package main

import (
	"strings"
)

// Slices s into two substrings before and after the first occurrence of sep.
func splitOnce(s, sep string) (string, string) {
	pos := strings.Index(s, sep)
	if pos == -1 {
		return s, ""
	}
	return s[:pos], s[pos+len(sep):]
}

// Parses environment into a key-value map
func mapEnviron(envv []string) map[string]string {
	vars := map[string]string{}

	for _, env := range envv {
		key, value := splitOnce(env, "=")
		vars[key] = value
	}

	return vars
}
