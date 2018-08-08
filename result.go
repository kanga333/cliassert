package main

import (
	"fmt"
	"strings"
)

type result struct {
	returnCode int
	stdout     string
	stderr     string
}

func (r *result) assertReturnCode(exprct int) bool {
	return r.returnCode == exprct
}

func (r *result) assertStdout(c assertCase) bool {
	return c.assert(r.stdout)
}

func (r *result) assertStderr(c assertCase) bool {
	return c.assert(r.stderr)
}

func (r *result) show() string {
	code := fmt.Sprintf("Exit code: %d", r.returnCode)
	stdout := fmt.Sprintf("Stdout: %s", r.stdout)
	stderr := fmt.Sprintf("Stderr: %s", r.stderr)
	result := []string{code, stdout, stderr}

	return fmt.Sprintln(strings.Join(result, "\n"))
}
