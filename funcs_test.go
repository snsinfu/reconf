package main

import (
	"strings"
	"testing"
	"text/template"
)

func render(src string, input interface{}) (string, error) {
	tmpl := template.New("test")
	tmpl.Funcs(templateFuncs)

	if _, err := tmpl.Parse(src); err != nil {
		return "", err
	}

	out := &strings.Builder{}
	if err := tmpl.Execute(out, input); err != nil {
		return "", err
	}

	return out.String(), nil
}

type funcTestCase struct {
	src    string
	input  interface{}
	expect string
}

type funcTestCases []funcTestCase

func (testCases funcTestCases) Run(t *testing.T) {
	t.Helper()

	for _, testCase := range testCases {
		actual, err := render(testCase.src, testCase.input)

		if err != nil {
			t.Errorf("unexepcted error: %v\n%#v", err, testCase)
			continue
		}
		if actual != testCase.expect {
			t.Errorf("incorrect result: %q\n%#v", actual, testCase)
		}
	}
}

func Test_templateFuncs_split(t *testing.T) {
	testCases := funcTestCases{
		{
			`{{ $ | split ":" }}`,
			"",
			"[]",
		},
		{
			`{{ $ | split ":" }}`,
			"quick",
			"[quick]",
		},
		{
			`{{ $ | split ":" }}`,
			"quick:brown:fox:jumps",
			"[quick brown fox jumps]",
		},
		{
			`{{ $ | split " <> " }}`,
			"quick <> brown <> fox <> jumps",
			"[quick brown fox jumps]",
		},
	}
	testCases.Run(t)
}

func Test_templateFuncs_join(t *testing.T) {
	testCases := funcTestCases{
		{
			`{{ $ | join ":" }}`,
			[]string{},
			"",
		},
		{
			`{{ $ | join ":" }}`,
			[]string{"quick"},
			"quick",
		},
		{
			`{{ $ | join ":" }}`,
			[]string{"quick", "brown", "fox", "jumps"},
			"quick:brown:fox:jumps",
		},
		{
			`{{ $ | join " <> " }}`,
			[]string{"quick", "brown", "fox", "jumps"},
			"quick <> brown <> fox <> jumps",
		},
	}
	testCases.Run(t)
}

func Test_templateFuncs_before(t *testing.T) {
	testCases := funcTestCases{
		{
			`{{ $ | before ":" }}`,
			"",
			"",
		},
		{
			`{{ $ | before ":" }}`,
			"quick",
			"quick",
		},
		{
			`{{ $ | before ":" }}`,
			"quick:brown:fox:jumps",
			"quick",
		},
		{
			`{{ $ | before " <> " }}`,
			"quick <> brown <> fox <> jumps",
			"quick",
		},
	}
	testCases.Run(t)
}

func Test_templateFuncs_after(t *testing.T) {
	testCases := funcTestCases{
		{
			`{{ $ | after ":" }}`,
			"",
			"",
		},
		{
			`{{ $ | after ":" }}`,
			"quick",
			"",
		},
		{
			`{{ $ | after ":" }}`,
			"quick:brown:fox:jumps",
			"brown:fox:jumps",
		},
		{
			`{{ $ | after " <> " }}`,
			"quick <> brown <> fox <> jumps",
			"brown <> fox <> jumps",
		},
	}
	testCases.Run(t)
}

func Test_tempateFuncs_nonempty(t *testing.T) {
	testCases := funcTestCases{
		{
			`{{ $ | nonempty }}`,
			[]string{},
			"[]",
		},
		{
			`{{ $ | nonempty }}`,
			[]string{"", ""},
			"[]",
		},
		{
			`{{ $ | nonempty }}`,
			[]string{"quick", "", "brown", "", "fox"},
			"[quick brown fox]",
		},
	}
	testCases.Run(t)
}

func Test_tempateFuncs_strip(t *testing.T) {
	testCases := funcTestCases{
		// Strings
		{
			`{{ $ | strip }}`,
			"",
			"",
		},
		{
			`{{ $ | strip }}`,
			" ",
			"",
		},
		{
			`{{ $ | strip }}`,
			"  quick brown fox  ",
			"quick brown fox",
		},
		{
			`{{ $ | strip }}`,
			"\t\nquick brown fox\n\n",
			"quick brown fox",
		},

		// Arrays
		{
			`{{ $ | strip }}`,
			[]string{},
			"[]",
		},
		{
			`{{ $ | strip }}`,
			[]string{" quick ", " brown ", " fox "},
			"[quick brown fox]",
		},
	}
	testCases.Run(t)
}
