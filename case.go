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

func newContainCase(substr string, isNot bool) *containCase {
	return &containCase{
		substr: substr,
		isNot:  isNot,
	}
}

func (c *containCase) assert(input string) bool {
	ok := strings.Contains(input, c.substr)
	if c.isNot {
		return !ok
	}
	return ok
}

type regexCase struct {
	pattern *regexp.Regexp
	isNot   bool
}

func newRegexCase(pattern string, isNot bool) (*regexCase, error) {
	r, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	return &regexCase{
		pattern: r,
		isNot:   isNot,
	}, nil
}

func (c *regexCase) assert(input string) bool {
	ok := c.pattern.MatchString(input)
	if c.isNot {
		return !ok
	}
	return ok
}
