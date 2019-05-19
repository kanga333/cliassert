package main

import (
	"os"

	"github.com/kanga333/cliassert"
)

func main() {
	os.Exit((&cliassert.CLI{ErrStream: os.Stderr, OutStream: os.Stdout, InStream: os.Stdin}).Run(os.Args[1:]))
}
