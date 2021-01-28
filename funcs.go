package main

import (
	"errors"
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

	// Removes empty strings from an array of strings.
	"nonempty": func(arr []string) []string {
		r := []string{}
		for _, s := range arr {
			if s != "" {
				r = append(r, s)
			}
		}
		return r
	},

	// Removes all leading and trailing white space of string. If an array of
	// string is given, it works on all the elements.
	"strip": func(arg interface{}) (interface{}, error) {
		switch arg := arg.(type) {
		case string:
			return strings.TrimSpace(arg), nil
		case []string:
			r := []string{}
			for _, s := range arg {
				r = append(r, strings.TrimSpace(s))
			}
			return r, nil
		}
		return nil, errors.New("strip expects string or array of strings as an input")
	},
}
