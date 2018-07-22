package main

type result struct {
	returnCode int
	stdout     string
	stderr     string
}

func (r *result) assertReturnCode(exprct int) bool {
	return r.returnCode == exprct
}

func (r *result) assertStdout(t testCase) (bool, string) {
	return t.assert(r.stdout)
}

func (r *result) assertStderr(t testCase) (bool, string) {
	return t.assert(r.stderr)
}
