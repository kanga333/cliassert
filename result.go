package main

type result struct {
	returnCode int
	stdout     string
	stderr     string
}

func (r *result) assertReturnCode(exprct int) bool {
	return r.returnCode == exprct
}

func (r *result) assertStdout(c assertCase) bool {
	return c.assert(r.stdout)
}

func (r *result) assertStderr(c assertCase) bool {
	return c.assert(r.stderr)
}
