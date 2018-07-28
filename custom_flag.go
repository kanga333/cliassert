package main

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
	generate() []assertCase
}

type containFlag struct {
	stringSlice
}

func (f *containFlag) generate() []assertCase {
	return generateAssertCase(f.stringSlice, newContainCase)
}

type notContainFlag struct {
	stringSlice
}

func (f *notContainFlag) generate() []assertCase {
	return generateAssertCase(f.stringSlice, newNotContainCase)
}

type regexFlag struct {
	stringSlice
}

func (f *regexFlag) generate() []assertCase {
	return generateAssertCase(f.stringSlice, newRegexCase)
}

type notRegexFlag struct {
	stringSlice
}

func (f *notRegexFlag) generate() []assertCase {
	return generateAssertCase(f.stringSlice, newNotRegexCase)
}

func generateAssertCase(ss stringSlice, f func(string) assertCase) []assertCase {
	var cases []assertCase
	for _, s := range ss {
		c := f(s)
		cases = append(cases, c)
	}
	return cases
}
