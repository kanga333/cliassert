package main

type testCase interface {
	assert(input string) bool
}
