package cliassert

import (
	"reflect"
	"testing"
)

func TestAssertion_assert(t *testing.T) {
	mock := []AssertCase{
		&mockCase{},
	}
	cases := []struct {
		Name      string
		Assertion Assertion
		Want1     bool
		Want2     *Result
	}{
		{
			"ok",
			Assertion{mock, mock, mock, "ok", "ok", "ok"},
			true,
			&Result{"ok", "ok", "ok", []string{"exit-status mock", "stdout mock", "stderr mock"}, nil},
		},
		{
			"ng_exitStatus",
			Assertion{mock, mock, mock, "ng", "ok", "ok"},
			false,
			&Result{"ng", "ok", "ok", []string{"stdout mock", "stderr mock"}, []string{"exit-status mock"}},
		},
		{
			"ng_stdout",
			Assertion{mock, mock, mock, "ok", "ng", "ok"},
			false,
			&Result{"ok", "ng", "ok", []string{"exit-status mock", "stderr mock"}, []string{"stdout mock"}},
		},
		{
			"ng_stderr",
			Assertion{mock, mock, mock, "ok", "ok", "ng"},
			false,
			&Result{"ok", "ok", "ng", []string{"exit-status mock", "stdout mock"}, []string{"stderr mock"}},
		},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got1, got2 := c.Assertion.Assert()
			if got1 != c.Want1 {
				t.Errorf("assert got:%v,want:%v:", got1, c.Want1)
			}
			if !reflect.DeepEqual(got2, c.Want2) {
				t.Errorf("assert got:\n%v,want:\n%v:", got2, c.Want2)
			}
		})
	}
}
