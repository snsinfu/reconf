package main

import (
	"os"
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

func Test_environ(t *testing.T) {
	key := "TEST_ENV_VAR"
	expect := "96777b30-74a7-4c09-a13f-7dc2e11ebdaa"

	if err := os.Setenv(key, expect); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	env := environ()

	actual, ok := env[key]
	if !ok {
		t.Errorf("key %q is not mapped", key)
	}
	if actual != expect {
		t.Errorf("key %q is incorrect: %q, want %q", key, actual, expect)
	}
}
