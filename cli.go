package cliassert

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var (
	verbose        *bool
	pass           *bool
	exitStatus     equalCaseFlag
	stdout         containFlag
	stderr         containFlag
	stdoutRegex    regexFlag
	stderrRegex    regexFlag
	notStdout      notContainFlag
	notStderr      notContainFlag
	notStdoutRegex notRegexFlag
	notStderrRegex notRegexFlag
)

const (
	exitStatusOK = iota
	exitStatusAssertFailure
	exitStatusError
)

func flagParse(args []string) ([]string, error) {
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	verbose = f.Bool("v", false, "Show verbose")
	pass = f.Bool("pass", false, "Pass stdout")
	f.Var(&exitStatus, "exit-status", "String equal to exit-status")
	f.Var(&stdout, "stdout-contain", "String contained in stdout")
	f.Var(&stderr, "stderr-contain", "String contained in stderr")
	f.Var(&stdoutRegex, "stdout-match", "Regex matching stdout")
	f.Var(&stderrRegex, "stderr-match", "Regex matching stderr")
	f.Var(&notStdout, "stdout-not-contain", "String not contained in stdout")
	f.Var(&notStderr, "stderr-not-contain", "String not contained in stderr")
	f.Var(&notStdoutRegex, "stdout-not-match", "Regex not matching stdout")
	f.Var(&notStderrRegex, "stderr-not-match", "Regex not matching stderr")
	err := f.Parse(args)
	if err != nil {
		return nil, err
	}
	return f.Args(), nil
}

// CLI is the struct that handles cli application.
type CLI struct {
	InStream             *os.File
	OutStream, ErrStream io.Writer
}

// Run the cliassert.
func (c *CLI) Run(args []string) int {
	parsedArgs, err := flagParse(args)
	if err != nil {
		fmt.Fprintf(c.ErrStream, "Fail to parse flag: %v\n", err)
		return exitStatusError
	}

	assertion, err := c.buildAssertion(parsedArgs)
	if err != nil {
		fmt.Fprintf(c.ErrStream, "Fail to build assertion case: %v\n", err)
		return exitStatusError
	}

	ok, result := assertion.Assert()
	output, err := renderResult(*verbose, result)
	if err != nil {
		fmt.Fprintf(c.ErrStream, "Fail to render output: %v\n", err)
		return exitStatusError
	}
	fmt.Fprint(c.ErrStream, output)

	if *pass {
		fmt.Fprint(c.OutStream, result.Stdout())
	}

	if !ok {
		return exitStatusAssertFailure
	}
	return exitStatusOK
}

func (c *CLI) buildAssertion(args []string) (*Assertion, error) {
	builder := AssertionBuilder{}

	builder.AppendExitStatusCases(exitStatus.Build())

	builder.AppendStdoutCases(stdout.Build())
	builder.AppendStdoutCases(stdoutRegex.Build())
	builder.AppendStdoutCases(notStdout.Build())
	builder.AppendStdoutCases(notStdoutRegex.Build())

	builder.AppendStderrCases(stderr.Build())
	builder.AppendStderrCases(stderrRegex.Build())
	builder.AppendStderrCases(notStderr.Build())
	builder.AppendStderrCases(notStderrRegex.Build())

	if c.stdinPiped() {
		if len(args) != 0 {
			return nil, errors.New("pipe and command can not be used simultaneously")
		}

		stdin, err := ioutil.ReadAll(c.InStream)
		if err != nil {
			return nil, err
		}
		return builder.BuildWithStdin(string(stdin))
	}
	return builder.BuildWithCommand(args)
}

func (c *CLI) stdinPiped() bool {
	fi, err := c.InStream.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeNamedPipe != 0
}

func renderResult(verbose bool, result *Result) (string, error) {
	if verbose {
		return result.ShowDetails()
	}
	return result.Show()
}
