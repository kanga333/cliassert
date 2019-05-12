package main

import (
	"os"

	"github.com/kanga333/cliassert"
)

func main() {
	os.Exit((&cliassert.CLI{ErrStream: os.Stderr, OutStream: os.Stdout}).Run(os.Args[1:]))
}
