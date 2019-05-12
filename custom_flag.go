package cliassert

import "fmt"

type stringSlice []string

func (s *stringSlice) String() string {
	return fmt.Sprintf("%v", *s)
}

func (s *stringSlice) Set(v string) error {
	*s = append(*s, v)
	return nil
}

type assertCaseFlag interface {
	Build() []AssertCase
}

type containFlag struct {
	stringSlice
}

func (f *containFlag) Build() []AssertCase {
	return buildAssertCase(f.stringSlice, NewContainCase)
}

type notContainFlag struct {
	stringSlice
}

func (f *notContainFlag) Build() []AssertCase {
	return buildAssertCase(f.stringSlice, NewNotContainCase)
}

type regexFlag struct {
	stringSlice
}

func (f *regexFlag) Build() []AssertCase {
	return buildAssertCase(f.stringSlice, NewRegexCase)
}

type notRegexFlag struct {
	stringSlice
}

func (f *notRegexFlag) Build() []AssertCase {
	return buildAssertCase(f.stringSlice, NewNotRegexCase)
}

type equalCaseFlag struct {
	stringSlice
}

func (f *equalCaseFlag) Build() []AssertCase {
	return buildAssertCase(f.stringSlice, NewEqualCase)
}

func buildAssertCase(ss stringSlice, f func(string) AssertCase) []AssertCase {
	var cases []AssertCase
	for _, s := range ss {
		c := f(s)
		cases = append(cases, c)
	}
	return cases
}
