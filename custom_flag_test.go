package cliassert

import (
	"reflect"
	"testing"
)

func TestAssertCaseFlag(t *testing.T) {
	cases := []struct {
		Name string
		Flag assertCaseFlag
		Want AssertCase
	}{
		{"containFlag", &containFlag{stringSlice{"1"}}, &ContainCase{}},
		{"notContainFlag", &notContainFlag{stringSlice{"1"}}, &NotContainCase{}},
		{"regexFlag", &regexFlag{stringSlice{"1"}}, &RegexCase{}},
		{"notRegexFlag", &notRegexFlag{stringSlice{"1"}}, &NotRegexCase{}},
		{"equalCaseFlag", &equalCaseFlag{stringSlice{"1"}}, &EqualCase{}},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := c.Flag.Build()
			if reflect.TypeOf(got[0]) != reflect.TypeOf(c.Want) {
				t.Errorf("Build type got: %v, want: %v", reflect.TypeOf(got[0]), reflect.TypeOf(c.Want))
			}
		})
	}
}
