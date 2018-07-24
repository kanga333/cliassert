package main

import (
	"regexp"
	"strings"
)

type assertCase interface {
	assert(input string) bool
}

type containCase struct {
	substr string
	isNot  bool
}

func (c *containCase) assert(input string) bool {
	ok := strings.Contains(input, c.substr)
	if c.isNot {
		return !ok
	}
	return ok
}

type regexCase struct {
	patern regexp.Regexp
	isNot  bool
}

func (c *regexCase) assert(input string) bool {
	ok := c.patern.MatchString(input)
	if c.isNot {
		return !ok
	}
	return ok
}
