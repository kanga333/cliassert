package main

import (
	"fmt"
	"regexp"
	"strings"
)

type assertCase interface {
	assert(input string) bool
	describe() string
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

func (c *containCase) describe() string {
	if c.isNot {
		msg := fmt.Sprintf("should not contain %s\n", c.substr)
		return msg
	}
	msg := fmt.Sprintf("should contain %s\n", c.substr)
	return msg
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

func (c *regexCase) describe() string {
	if c.isNot {
		msg := fmt.Sprintf("should not match %s\n", c.pattern)
		return msg
	}
	msg := fmt.Sprintf("should match %s\n", c.pattern)
	return msg
}
