package cliassert

import (
	"reflect"
	"testing"
)

func TestAssertionBuilder(t *testing.T) {
	mock := []AssertCase{
		&mockCase{},
	}
	cases := []struct {
		Name            string
		ExitStatusCases []AssertCase
		StdoutCases     []AssertCase
		StderrCases     []AssertCase
		Want            AssertionBuilder
	}{
		{"append", mock, mock, mock, AssertionBuilder{mock, mock, mock}},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := AssertionBuilder{}
			got.AppendExitStatusCases(c.ExitStatusCases)
			got.AppendStdoutCases(c.StdoutCases)
			got.AppendStderrCases(c.StderrCases)
			if !reflect.DeepEqual(got, c.Want) {
				t.Errorf("Append \ngot:\n%v,\nwant:\n%v:", got, c.Want)
			}

		})
	}

}

func TestAssertionBuilder_BuildWithCommand(t *testing.T) {
	cases := []struct {
		Name string
		Cmd  []string
		Want *Assertion
	}{
		{"one_cmd", []string{"echo"},
			&Assertion{
				exitStatus: "0",
				stdout:     "\n",
				stderr:     "",
			},
		},
		{"stdout", []string{"echo", "stdout"},
			&Assertion{
				exitStatus: "0",
				stdout:     "stdout\n",
				stderr:     "",
			},
		},
		{"stderr", []string{"sh", "-c", "echo 1>&2 stderr"},
			&Assertion{
				exitStatus: "0",
				stdout:     "",
				stderr:     "stderr\n",
			},
		},
		{"fail", []string{"cat", "nonexistent_file"},
			&Assertion{
				exitStatus: "1",
				stdout:     "",
				stderr:     "cat: nonexistent_file: No such file or directory\n",
			},
		},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			a := AssertionBuilder{}
			got, err := a.BuildWithCommand(c.Cmd[0], c.Cmd[1:]...)
			if err != nil {
				t.Fatalf("BuildWithCommand return error: %v", err)
			}
			if !reflect.DeepEqual(got, c.Want) {
				t.Errorf("BuildWithCommand \ngot:\n%v,\nwant:\n%v:", got, c.Want)
			}

		})
	}
}
