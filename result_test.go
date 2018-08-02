package main

import (
	"testing"
)

func TestResult(t *testing.T) {
	cases := [][]struct {
		subString string
		input     string
		want      bool
	}{
		{
			{subString: "stdout", want: true},
			{subString: "stderr", want: false},
		},
		{
			{subString: "stdout", want: false},
			{subString: "stderr", want: true},
		},
	}

	r := &result{
		returnCode: 0,
		stdout:     "stdout",
		stderr:     "stderr",
	}

	got := r.assertReturnCode(0)
	if got != true {
		t.Errorf("assertReturnCode  got: %v, want: %v:", got, true)
	}

	for _, c := range cases[0] {
		cc := newContainCase(c.subString)
		got := r.assertStdout(cc)
		if got != c.want {
			t.Errorf("Assert %s %s got: %v, want: %v:", r.stdout, cc.describe(), got, c.want)
		}
	}

	for _, c := range cases[1] {
		cc := newContainCase(c.subString)
		got := r.assertStderr(cc)
		if got != c.want {
			t.Errorf("Assert %s %s got: %v, want: %v:", r.stderr, cc.describe(), got, c.want)
		}
	}
}
