package main

import (
	"testing"
)

func TestAssertion(t *testing.T) {
	a := newAssertion(0)
	stdout := newContainCase("stdout")
	stderr := newContainCase("stderr")
	a.appendStdoutCase(stdout)
	a.appendStderrCase(stderr)

	r := result{
		returnCode: 0,
		stdout:     "stdout",
		stderr:     "stderr",
	}

	got, _ := a.assertCliResult(r)
	if got != true {
		t.Errorf("assertCliResult got: %v, want: %v:", got, true)
	}
}

func TestAssertionCodeFail(t *testing.T) {
	a := newAssertion(0)
	r := result{
		returnCode: 1,
	}

	got, gotMsg := a.assertCliResult(r)
	if got != false {
		t.Errorf("assertCliResult got: %v, want: %v:", got, false)
	}

	if want := "return code should be 0"; gotMsg != want {
		t.Errorf("assertCliResult got: %v, want: %v:", gotMsg, want)
	}
}

func TestAssertionCaseFail(t *testing.T) {
	a := newAssertion(0)
	stdout := newContainCase("fail")
	stderr := newContainCase("stderr")
	a.appendStdoutCase(stdout)
	a.appendStderrCase(stderr)

	r := result{
		returnCode: 0,
		stdout:     "stdout",
		stderr:     "stderr",
	}

	got, gotMsg := a.assertCliResult(r)
	if got != false {
		t.Errorf("assertCliResult got: %v, want: %v:", got, false)
	}

	if want := stdout.describe(); gotMsg != want {
		t.Errorf("assertCliResult got: %v, want: %v:", gotMsg, want)
	}
}
