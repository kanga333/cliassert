package cliassert

import (
	"bytes"
	"errors"
	"os/exec"
	"strconv"
	"syscall"
)

// AssertionBuilder is the builder that build assertion object.
// Add a test case to this object to build an assertion.
type AssertionBuilder struct {
	exitStatusCases []AssertCase
	stdoutCases     []AssertCase
	stderrCases     []AssertCase
}

// AppendExitStatusCases append assertion cases to check the exit status code.
func (a *AssertionBuilder) AppendExitStatusCases(cases []AssertCase) {
	a.exitStatusCases = append(a.exitStatusCases, cases...)
}

// AppendStdoutCases append assertion cases to check the stdout.
func (a *AssertionBuilder) AppendStdoutCases(cases []AssertCase) {
	a.stdoutCases = append(a.stdoutCases, cases...)
}

// AppendStderrCases append assertion cases to check the stderr.
func (a *AssertionBuilder) AppendStderrCases(cases []AssertCase) {
	a.stderrCases = append(a.stderrCases, cases...)
}

// BuildWithCommand built assertion form command　result.
func (a *AssertionBuilder) BuildWithCommand(args []string) (*Assertion, error) {
	if len(args) == 0 {
		return nil, errors.New("missing argument for command")
	}

	var stdout, stderr bytes.Buffer
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	exitStatus := 0

	err := cmd.Run()
	if err != nil {
		exitError, ok := err.(*exec.ExitError)
		if !ok {
			return nil, err
		}
		ws := exitError.Sys().(syscall.WaitStatus)
		exitStatus = ws.ExitStatus()
	}

	return &Assertion{
		exitStatusCases: a.exitStatusCases,
		stdoutCases:     a.stdoutCases,
		stderrCases:     a.stderrCases,
		exitStatus:      strconv.Itoa(exitStatus),
		stdout:          stdout.String(),
		stderr:          stderr.String(),
	}, nil
}

// BuildWithStdin built assertion form stdin.
func (a *AssertionBuilder) BuildWithStdin(stdin string) (*Assertion, error) {
	if (len(a.exitStatusCases) != 0) || (len(a.stderrCases) != 0) {
		return nil, errors.New("pipe can only be used for stdout testing")
	}
	return &Assertion{
		stdoutCases: a.stdoutCases,
		stdout:      stdin,
	}, nil
}
