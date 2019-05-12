package cliassert

import (
	"fmt"
	"regexp"
	"strings"
)

// AssertCase represents one assertion case.
// It can examine and describe the assertion.
type AssertCase interface {
	Assert(input string) bool
	Describe() string
}

// ContainCase represents one assertion case.
// It checks if the input contains strings.
type ContainCase struct {
	substr string
}

// NewContainCase creates a case to inspect a given string.
func NewContainCase(substr string) AssertCase {
	return &ContainCase{substr: substr}
}

// Assert inspects case assertion.
func (c *ContainCase) Assert(input string) bool {
	return strings.Contains(input, c.substr)
}

// Describe Describes the case.
func (c *ContainCase) Describe() string {
	return fmt.Sprintf("should contain %s", c.substr)
}

// NotContainCase represents one assertion case.
// It checks that the input does not contain strings.
type NotContainCase struct {
	substr string
}

// NewNotContainCase creates a case to inspect a given string.
func NewNotContainCase(substr string) AssertCase {
	return &NotContainCase{substr: substr}
}

// Assert inspects case assertion.
func (c *NotContainCase) Assert(input string) bool {
	return !strings.Contains(input, c.substr)
}

// Describe Describes the case.
func (c *NotContainCase) Describe() string {
	return fmt.Sprintf("should not contain %s", c.substr)
}

// RegexCase represents one assertion case.
// It checks if the input match regex pattern.
type RegexCase struct {
	pattern *regexp.Regexp
}

// NewRegexCase creates a case to inspect a given string.
func NewRegexCase(pattern string) AssertCase {
	r := regexp.MustCompile(pattern)
	return &RegexCase{r}
}

// Assert inspects case assertion.
func (c *RegexCase) Assert(input string) bool {
	return c.pattern.MatchString(input)
}

// Describe Describes the case.
func (c *RegexCase) Describe() string {
	return fmt.Sprintf("should match %s", c.pattern)
}

// NotRegexCase represents one assertion case.
// It checks that the input does not match regex pattern.
type NotRegexCase struct {
	pattern *regexp.Regexp
}

// NewNotRegexCase creates a case to inspect a given string.
func NewNotRegexCase(pattern string) AssertCase {
	r := regexp.MustCompile(pattern)
	return &NotRegexCase{r}
}

// Assert inspects case assertion.
func (c *NotRegexCase) Assert(input string) bool {
	return !c.pattern.MatchString(input)
}

// Describe Describes the case.
func (c *NotRegexCase) Describe() string {
	return fmt.Sprintf("should not match %s", c.pattern)
}

// EqualCase represents one assertion case.
// It checks that the input and the string are the same.
type EqualCase struct {
	want string
}

// NewEqualCase creates a case to inspect a given string.
func NewEqualCase(want string) AssertCase {
	return &EqualCase{want}
}

// Assert inspects case assertion.
func (c *EqualCase) Assert(input string) bool {
	c.Describe()
	return c.want == input
}

// Describe Describes the case.
func (c *EqualCase) Describe() string {
	return fmt.Sprintf("should be equal %s", c.want)
}
