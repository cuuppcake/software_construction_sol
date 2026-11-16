package main

import (
	"fmt"
	"os"
	"pos/internal/bootstrap"
)

func main() {
	if err := bootstrap.Run(os.Args[1:], os.Stdin, os.Stdout); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
