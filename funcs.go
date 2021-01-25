package main

import (
	"strings"
	"text/template"
)

// Defines utility functions for use in templates.
var templateFuncs = template.FuncMap{
	"split": func(sep, s string) []string {
		return strings.Split(s, sep)
	},

	"join": func(sep string, elems []string) string {
		return strings.Join(elems, sep)
	},

	"before": func(sep, s string) string {
		sub, _ := splitOnce(s, sep)
		return sub
	},

	"after": func(sep, s string) string {
		_, sub := splitOnce(s, sep)
		return sub
	},
}
