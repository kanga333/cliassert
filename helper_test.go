package cliassert

import "flag"

var update = flag.Bool("update", false, "update golden files")

type mockCase struct{}

func (m *mockCase) Assert(Input string) bool {
	return Input == "ok"
}

func (m *mockCase) Describe() string {
	return "mock"
}
