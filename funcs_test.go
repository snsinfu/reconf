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
	src := `{{ $ | split ":" }}`
	input := "quick:brown:fox:jumps"
	expect := "[quick brown fox jumps]"

	actual, err := render(src, input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if actual != expect {
		t.Fatalf("incorrect: %q, want %q", actual, expect)
	}
}

func Test_templateFuncs_join(t *testing.T) {
	src := `{{ $ | join ":" }}`
	input := []string{"quick", "brown", "fox", "jumps"}
	expect := "quick:brown:fox:jumps"

	actual, err := render(src, input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if actual != expect {
		t.Fatalf("incorrect: %q, want %q", actual, expect)
	}
}

func Test_templateFuncs_before(t *testing.T) {
	src := `{{ $ | before ":" }}`
	input := "quick:brown:fox:jumps"
	expect := "quick"

	actual, err := render(src, input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if actual != expect {
		t.Fatalf("incorrect: %q, want %q", actual, expect)
	}
}

func Test_templateFuncs_after(t *testing.T) {
	src := `{{ $ | after ":" }}`
	input := "quick:brown:fox:jumps"
	expect := "brown:fox:jumps"

	actual, err := render(src, input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if actual != expect {
		t.Fatalf("incorrect: %q, want %q", actual, expect)
	}
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
