package main

import (
	"testing"
)

func TestContainFlag(t *testing.T) {
	cases := []struct {
		subString string
		input     string
		want      bool
	}{
		{subString: "t", input: "test", want: true},
		{subString: "n", input: "test", want: false},
	}

	var flag containFlag
	for _, c := range cases {
		flag.Set(c.subString)
	}

	assertCases := flag.generate()

	for i, ac := range assertCases {
		got := ac.assert(cases[i].input)
		if got != cases[i].want {
			t.Errorf("Assert %s %s got: %v, want: %v:", cases[i].input, ac.describe(), got, cases[i].want)
		}
	}
}

func TestNotContainFlag(t *testing.T) {
	cases := []struct {
		subString string
		input     string
		want      bool
	}{
		{subString: "t", input: "test", want: false},
		{subString: "n", input: "test", want: true},
	}

	var flag notContainFlag
	for _, c := range cases {
		flag.Set(c.subString)
	}

	assertCases := flag.generate()

	for i, ac := range assertCases {
		got := ac.assert(cases[i].input)
		if got != cases[i].want {
			t.Errorf("Assert %s %s got: %v, want: %v:", cases[i].input, ac.describe(), got, cases[i].want)
		}
	}
}

func TestRegexFlag(t *testing.T) {
	cases := []struct {
		subString string
		input     string
		want      bool
	}{
		{subString: "te.t", input: "test", want: true},
		{subString: "te..t", input: "test", want: false},
	}

	var flag regexFlag
	for _, c := range cases {
		flag.Set(c.subString)
	}

	assertCases := flag.generate()

	for i, ac := range assertCases {
		got := ac.assert(cases[i].input)
		if got != cases[i].want {
			t.Errorf("Assert %s %s got: %v, want: %v:", cases[i].input, ac.describe(), got, cases[i].want)
		}
	}
}

func TestNotRegexFlag(t *testing.T) {
	cases := []struct {
		subString string
		input     string
		want      bool
	}{
		{subString: "te.t", input: "test", want: false},
		{subString: "te..t", input: "test", want: true},
	}

	var flag notRegexFlag
	for _, c := range cases {
		flag.Set(c.subString)
	}

	assertCases := flag.generate()

	for i, ac := range assertCases {
		got := ac.assert(cases[i].input)
		if got != cases[i].want {
			t.Errorf("Assert %s %s got: %v, want: %v:", cases[i].input, ac.describe(), got, cases[i].want)
		}
	}
}
