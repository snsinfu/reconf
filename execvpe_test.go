package main

import (
	"syscall"
	"testing"
)

func Test_execvpe_errorEmptyPath(t *testing.T) {
	// Search path is empty, so the command won't be found.
	argv := []string{"echo", ""}
	paths := []string{}
	envs := []string{}
	err := execvpe(argv[0], paths, argv, envs)

	if err != syscall.ENOENT {
		t.Fatalf("unexpected error: %v", err)
	}
}

func Test_execvpe_errorNotInSearchPath(t *testing.T) {
	// This command is assumed not to be found.
	argv := []string{"none", ""}
	paths := []string{"/aaa", "/bbb"}
	envs := []string{}
	err := execvpe(argv[0], paths, argv, envs)

	if err != syscall.ENOENT {
		t.Fatalf("unexpected error: %v", err)
	}
}

func Test_execvpe_errorExecNotFound(t *testing.T) {
	// This path is assumed to be non-existent.
	argv := []string{"/non/existing/executable", ""}
	paths := []string{}
	envs := []string{}
	err := execvpe(argv[0], paths, argv, envs)

	if err != syscall.ENOENT {
		t.Fatalf("unexpected error: %v", err)
	}
}
