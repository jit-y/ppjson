package main

import (
	"fmt"
	"os"

	"github.com/jit-y/ppjson"
)

func main() {
	printer := ppjson.NewPrinter(os.Stdin)
	j, err := printer.Pretty()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(j)
}
