package main

import (
	"os"
	"strings"
)

// Slices s into two substrings before and after the first occurrence of sep.
func splitOnce(s, sep string) (string, string) {
	pos := strings.Index(s, sep)
	if pos == -1 {
		return s, ""
	}
	return s[:pos], s[pos+1:]
}

// Returns a copy of the environment as a map.
func environ() map[string]string {
	vars := map[string]string{}

	for _, env := range os.Environ() {
		key, value := splitOnce(env, "=")
		vars[key] = value
	}

	return vars
}
