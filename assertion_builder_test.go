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
			got, err := a.BuildWithCommand(c.Cmd)
			if err != nil {
				t.Fatalf("BuildWithCommand return error: %v", err)
			}
			if !reflect.DeepEqual(got, c.Want) {
				t.Errorf("BuildWithCommand \ngot:\n%v,\nwant:\n%v:", got, c.Want)
			}

		})
	}
}

func TestAssertionBuilder_BuildWithCommand_Error(t *testing.T) {
	a := AssertionBuilder{}
	got, err := a.BuildWithCommand([]string{})
	if err == nil {
		t.Fatalf("BuildWithCommand with empty arguments should return error but return: %v", got)
	}
}

func TestAssertionBuilder_BuildWithStdin(t *testing.T) {
	cases := []struct {
		Name  string
		Stdin string
		Want  *Assertion
	}{
		{"ok", "stdin",
			&Assertion{
				stdoutCases: []AssertCase{&mockCase{}},
				stdout:      "stdin",
			},
		},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			a := AssertionBuilder{}
			a.AppendStdoutCases([]AssertCase{&mockCase{}})
			got, err := a.BuildWithStdin(c.Stdin)
			if err != nil {
				t.Fatalf("BuildWithStdin return error: %v", err)
			}
			if !reflect.DeepEqual(got, c.Want) {
				t.Errorf("BuildWithStdin \ngot:\n%v,\nwant:\n%v:", got, c.Want)
			}

		})
	}
}

func TestAssertionBuilder_BuildWithStdin_Error(t *testing.T) {
	cases := []struct {
		Name    string
		Builder AssertionBuilder
	}{
		{"stderr cases", AssertionBuilder{
			stderrCases: []AssertCase{&mockCase{}},
		}},
		{"exit-status cases", AssertionBuilder{
			exitStatusCases: []AssertCase{&mockCase{}},
		}},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got, err := c.Builder.BuildWithStdin("")
			if err == nil {
				t.Fatalf("BuildWithStdin with %s should return error but return: %v", c.Name, got)
			}
		})
	}
}
