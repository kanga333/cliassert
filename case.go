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
}

func newContainCase(substr string) assertCase {
	return &containCase{substr: substr}
}

func (c *containCase) assert(input string) bool {
	return strings.Contains(input, c.substr)
}

func (c *containCase) describe() string {
	return fmt.Sprintf("should contain %s\n", c.substr)
}

type notContainCase struct {
	substr string
}

func newNotContainCase(substr string) assertCase {
	return &notContainCase{substr: substr}
}

func (c *notContainCase) assert(input string) bool {
	return !strings.Contains(input, c.substr)
}

func (c *notContainCase) describe() string {
	return fmt.Sprintf("should not contain %s\n", c.substr)
}

type regexCase struct {
	pattern *regexp.Regexp
}

func newRegexCase(pattern string) assertCase {
	r := regexp.MustCompile(pattern)
	return &regexCase{r}
}

func (c *regexCase) assert(input string) bool {
	return c.pattern.MatchString(input)
}

func (c *regexCase) describe() string {
	return fmt.Sprintf("should match %s\n", c.pattern)
}

type notRegexCase struct {
	pattern *regexp.Regexp
}

func newNotRegexCase(pattern string) assertCase {
	r := regexp.MustCompile(pattern)
	return &notRegexCase{r}
}

func (c *notRegexCase) assert(input string) bool {
	return !c.pattern.MatchString(input)
}

func (c *notRegexCase) describe() string {
	return fmt.Sprintf("should not match %s\n", c.pattern)
}
