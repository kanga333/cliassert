package main

import (
	"testing"
)

func TestContainCase(t *testing.T) {
	c := newContainCase("t")
	okCase := "test"
	ngCase := "ng"

	got := c.assert(okCase)
	if got != true {
		t.Errorf("Assert %s got: %v, want: true:", okCase, got)
	}

	got = c.assert(ngCase)
	if got != false {
		t.Errorf("Assert %s got: %v, want: false:", ngCase, got)
	}
}
