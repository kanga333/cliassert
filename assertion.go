package cliassert

import (
	"fmt"
)

// Assertion is the struct that assert the command execution result.
type Assertion struct {
	exitStatusCases []AssertCase
	stdoutCases     []AssertCase
	stderrCases     []AssertCase
	exitStatus      string
	stdout          string
	stderr          string
}

// Assert asserts various command execution results.
func (a *Assertion) Assert() (bool, *Result) {
	succeeded := true
	var successes, failures []string

	for _, c := range a.exitStatusCases {
		s := fmt.Sprintf("exit-status %s", c.Describe())
		if !c.Assert(a.exitStatus) {
			succeeded = false
			failures = append(failures, s)
			continue
		}
		successes = append(successes, s)
	}

	for _, c := range a.stdoutCases {
		s := fmt.Sprintf("stdout %s", c.Describe())
		if !c.Assert(a.stdout) {
			succeeded = false
			failures = append(failures, s)
			continue
		}
		successes = append(successes, s)
	}

	for _, c := range a.stderrCases {
		s := fmt.Sprintf("stderr %s", c.Describe())
		if !c.Assert(a.stderr) {
			succeeded = false
			failures = append(failures, s)
			continue
		}
		successes = append(successes, s)
	}

	return succeeded, &Result{
		exitStatus: a.exitStatus,
		stdout:     a.stdout,
		stderr:     a.stderr,
		successes:  successes,
		failures:   failures,
	}
}
