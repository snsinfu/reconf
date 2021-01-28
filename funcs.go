package main

import (
	"strings"
	"text/template"
)

// Defines utility functions for use in templates.
var templateFuncs = template.FuncMap{
	// These functions are meant to be called in a pipe chain, where
	//   $input | split ":" | nonempty | joinn ":"
	// is equivalent to
	//   join (":" nonempty (split ":" $input))
	// Hence, the functions take arguments in unusual order.

	// Splits `s` into all substrings separated by `sep`.
	"split": func(sep, s string) []string {
		return strings.Split(s, sep)
	},

	// Concatenates `elems` into a single string with separator `sep`.
	"join": func(sep string, elems []string) string {
		return strings.Join(elems, sep)
	},

	// Returns the substring of `s` before `sep` (non-inclusive). It returns
	// the entire string `s` if `sep` is not in the string.
	"before": func(sep, s string) string {
		sub, _ := splitOnce(s, sep)
		return sub
	},

	// Returns the substring of `s` after `sep` (non-inclusive). It returns
	// the entire string `s` if `sep` is not in the string.
	"after": func(sep, s string) string {
		_, sub := splitOnce(s, sep)
		return sub
	},
}
