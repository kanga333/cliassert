package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	stdout         containFlag
	stderr         containFlag
	stdoutRegex    regexFlag
	stderrRegex    regexFlag
	notStdout      notContainFlag
	notStderr      notContainFlag
	notStdoutRegex notRegexFlag
	notStderrRegex notRegexFlag
)

func init() {
	flag.Var(&stdout, "stdout", "String contained in stdout")
	flag.Var(&stderr, "stderr", "String contained in stderr")
	flag.Var(&stdoutRegex, "stdout-regex", "Regex matching stdout")
	flag.Var(&stderrRegex, "stderr-regex", "Regex matching stderr")
	flag.Var(&notStdout, "not-stdout", "String not contained in stdout")
	flag.Var(&notStderr, "not-stderr", "String not contained in stderr")
	flag.Var(&notStdoutRegex, "not-stdout-regex", "Regex not matching stdout")
	flag.Var(&notStderrRegex, "not-stderr-regex", "Regex not matching stderr")
}

func main() {
	code := flag.Int("code", 0, "Expected return code")
	flag.Parse()
	args := flag.Args()

	assertion := newAssertion(*code)

	appendStdoutCase(assertion, &stdout)
	appendStdoutCase(assertion, &stdoutRegex)
	appendStdoutCase(assertion, &notStdout)
	appendStdoutCase(assertion, &notStdoutRegex)

	appendStderrCase(assertion, &stderr)
	appendStderrCase(assertion, &stdoutRegex)
	appendStderrCase(assertion, &notStdout)
	appendStderrCase(assertion, &notStdoutRegex)

	cmd := Command(args[0], args[1:]...)
	result := cmd.exec()

	ok, testResult := assertion.assertCliResult(*result)
	if !ok {
		fmt.Printf("fail: %s.", testResult)
		os.Exit(1)
	}
}

func appendStdoutCase(a *assertion, flags assertCaseFlag) {
	cases := flags.generate()
	for _, c := range cases {
		a.appendStdoutCase(c)
	}
}

func appendStderrCase(a *assertion, flags assertCaseFlag) {
	cases := flags.generate()
	for _, c := range cases {
		a.appendStderrCase(c)
	}
}
