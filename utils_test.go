package main

import (
	"reflect"
	"testing"
)

func Test_splitOnce(t *testing.T) {
	testCases := []struct {
		input string
		sep   string
		left  string
		right string
	}{
		{"a/b/c", "/", "a", "b/c"},
		{"a//b/c", "/", "a", "/b/c"},

		{"abc/", "/", "abc", ""},
		{"/abc", "/", "", "abc"},
		{"abc", "/", "abc", ""},

		{"/", "/", "", ""},
		{"", "/", "", ""},

		{"1::2", "::", "1", "2"},
	}

	for _, testCase := range testCases {
		left, right := splitOnce(testCase.input, testCase.sep)

		if left != testCase.left {
			t.Errorf("left is incorrect: %q, want %q", left, testCase.left)
		}
		if right != testCase.right {
			t.Errorf("right is incorrect: %q, want %q", right, testCase.right)
		}
	}
}

func Test_mapEnviron(t *testing.T) {
	testCases := []struct {
		envv   []string
		expect map[string]string
	}{
		{
			[]string{},
			map[string]string{},
		},
		{
			[]string{"A="},
			map[string]string{"A": ""},
		},
		{
			[]string{"A=1", "ABC=123"},
			map[string]string{"A": "1", "ABC": "123"},
		},
		{
			[]string{"A=a=1,b=2"},
			map[string]string{"A": "a=1,b=2"},
		},
	}

	for _, testCase := range testCases {
		actual := mapEnviron(testCase.envv)

		if !reflect.DeepEqual(actual, testCase.expect) {
			t.Errorf("incorrect result: %v, want %v", actual, testCase.expect)
		}
	}
}
