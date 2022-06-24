package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	w := os.Stdout
	if err := run(os.Args, w); err != nil {
		fmt.Fprintf(w, "%s\n", err)
		os.Exit(1)
	}
}

func run(args []string, w io.Writer) error {
	return nil
}
