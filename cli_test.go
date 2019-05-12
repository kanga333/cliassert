package cliassert

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func flagInit() {
	exitStatus = equalCaseFlag{}
	stdout = containFlag{}
	stderr = containFlag{}
	stdoutRegex = regexFlag{}
	stderrRegex = regexFlag{}
	notStdout = notContainFlag{}
	notStderr = notContainFlag{}
	notStdoutRegex = notRegexFlag{}
	notStderrRegex = notRegexFlag{}
}

func TestResult_CLI(t *testing.T) {
	cases := []struct {
		Name string
		Args []string
		Want int
	}{
		{
			"all_pass",
			[]string{
				"-v",
				"-exit-status", "0",
				"-stdout-contain", "stdout",
				"-stderr-contain", "stderr",
				"-stdout-match", ".*",
				"-stderr-match", ".*",
				"-stdout-not-contain", "ng",
				"-stderr-not-contain", "ng",
				"-stdout-not-match", "ng",
				"-stderr-not-match", "ng",
				"sh", "-c", "echo stdout ; echo 1>&2 stderr",
			},
			0,
		},
		{"success", []string{"-v", "-exit-status", "0", "echo", "stdout"}, 0},
		{"failure", []string{"-exit-status", "1", "echo", "stdout"}, 1},
		{"error", []string{"-exit-status", "1", "nonexistent_command"}, 2},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			var outStream, errStream bytes.Buffer
			cli := CLI{
				OutStream: &outStream,
				ErrStream: &errStream,
			}

			got := cli.Run(c.Args)
			defer flagInit()
			if got != c.Want {
				t.Errorf("Run got:%v, want:%v", got, c.Want)
			}

			gotStdout := outStream.Bytes()
			gotStderr := errStream.Bytes()
			goldenStdout := filepath.Join("fixtures/cli", c.Name+"_out.golden")
			goldenStderr := filepath.Join("fixtures/cli", c.Name+"_err.golden")
			if *update {
				ioutil.WriteFile(goldenStdout, gotStdout, 0644)
				ioutil.WriteFile(goldenStderr, gotStderr, 0644)
			}

			wantStdout, _ := ioutil.ReadFile(goldenStdout)
			if got, want := string(gotStdout), string(wantStdout); got != want {
				t.Errorf("Run stdout \ngot:\n%v,want:\n%v", got, want)
			}
			wantStderr, _ := ioutil.ReadFile(goldenStderr)
			if got, want := string(gotStderr), string(wantStderr); got != want {
				t.Errorf("Run stderr \ngot:\n%v,want:\n%v", got, want)
			}
		})
	}
}
