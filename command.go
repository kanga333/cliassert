package main

import (
	"bytes"
	"os/exec"
	"syscall"
)

type cmd struct {
	cmd *exec.Cmd
}

func Command(name string, arg ...string) *cmd {
	c := exec.Command(name, arg...)
	return &cmd{
		cmd: c,
	}
}

func (c *cmd) exec() *result {
	var stdout, stderr bytes.Buffer
	c.cmd.Stdout = &stdout
	c.cmd.Stderr = &stderr
	returnCode := 0

	err := c.cmd.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			returnCode = ws.ExitStatus()
		}
	}

	return &result{
		returnCode: returnCode,
		stdout:     stdout.String(),
		stderr:     stderr.String(),
	}
}
