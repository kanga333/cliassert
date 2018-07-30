package main

import (
	"testing"
)

func TestContainCase(t *testing.T) {
	cases := []struct {
		subString string
		input     string
		want      bool
	}{
		{subString: "t", input: "test", want: true},
		{subString: "n", input: "test", want: false},
	}

	for _, c := range cases {
		cc := newContainCase(c.subString)
		got := cc.assert(c.input)
		if got != c.want {
			t.Errorf("Assert %s %s got: %v, want: %v:", c.input, cc.describe(), got, c.want)
		}

	}
}

func TestNotContainCase(t *testing.T) {
	cases := []struct {
		subString string
		input     string
		want      bool
	}{
		{subString: "t", input: "test", want: false},
		{subString: "n", input: "test", want: true},
	}

	for _, c := range cases {
		cc := newNotContainCase(c.subString)
		got := cc.assert(c.input)
		if got != c.want {
			t.Errorf("Assert %s %s got: %v, want: %v:", c.input, cc.describe(), got, c.want)
		}

	}
}

func TestRegexCase(t *testing.T) {
	cases := []struct {
		subString string
		input     string
		want      bool
	}{
		{subString: "te.t", input: "test", want: true},
		{subString: "te..t", input: "test", want: false},
	}

	for _, c := range cases {
		cc := newRegexCase(c.subString)
		got := cc.assert(c.input)
		if got != c.want {
			t.Errorf("Assert %s %s got: %v, want: %v:", c.input, cc.describe(), got, c.want)
		}

	}
}

func TestNotRegexCase(t *testing.T) {
	cases := []struct {
		subString string
		input     string
		want      bool
	}{
		{subString: "te.t", input: "test", want: false},
		{subString: "te..t", input: "test", want: true},
	}

	for _, c := range cases {
		cc := newNotRegexCase(c.subString)
		got := cc.assert(c.input)
		if got != c.want {
			t.Errorf("Assert %s %s got: %v, want: %v:", c.input, cc.describe(), got, c.want)
		}

	}
}
