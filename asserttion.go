package main

import (
	"fmt"
)

type assertion struct {
	expectCode  int
	stdoutCases []assertCase
	stderrCases []assertCase
}

func newAssertion(code int) *assertion {
	a := &assertion{
		expectCode: code,
	}
	return a
}

func (a *assertion) appendStdoutCase(ac assertCase) {
	a.stdoutCases = append(a.stdoutCases, ac)
}

func (a *assertion) appendStderrCase(ac assertCase) {
	a.stderrCases = append(a.stderrCases, ac)
}

func (a *assertion) assertCliResult(r result) (bool, string) {
	if !r.assertReturnCode(a.expectCode) {
		return false, fmt.Sprint("return code should be ", a.expectCode)
	}

	for _, c := range a.stdoutCases {
		if !r.assertStdout(c) {
			desctibe := fmt.Sprintf("Stdout %s", c.describe())
			return false, desctibe
		}
	}

	for _, c := range a.stderrCases {
		if !r.assertStderr(c) {
			desctibe := fmt.Sprintf("Stderr %s", c.describe())
			return false, desctibe
		}
	}
	return true, ""
}
