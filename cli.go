package cliassert

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	verbose        *bool
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
	verbose = f.Bool("v", false, "show verbose")
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
	OutStream, ErrStream io.Writer
}

// Run the cliassert.
func (c *CLI) Run(args []string) int {
	parsedArgs, err := flagParse(args)
	if err != nil {
		fmt.Fprintf(c.ErrStream, "Flag parse error: %v\n", err)
		return exitStatusError
	}

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

	assertion, err := builder.BuildWithCommand(parsedArgs[0], parsedArgs[1:]...)
	if err != nil {
		fmt.Fprintf(c.ErrStream, "Command execution error: %v\n", err)
		return exitStatusError
	}

	ok, result := assertion.Assert()
	output, err := renderResult(*verbose, result)
	if err != nil {
		fmt.Fprintf(c.ErrStream, "Result rendering error: %v\n", err)
		return exitStatusError
	}
	fmt.Fprint(c.ErrStream, output)

	if !ok {
		return exitStatusAssertFailure
	}
	return exitStatusOK
}

func renderResult(verbose bool, result *Result) (string, error) {
	if verbose {
		return result.ShowDetails()
	}
	return result.Show()
}
