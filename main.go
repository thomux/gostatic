/*
CLI executable for the gostatic static page generator.
*/
package main

import (
	"os"
	"path/filepath"

	"github.com/thomux/gostatic/gostatic"
)

// main is the CLI executable and parses the command line options.
func main() {
	var dir string
	var err error

	if len(os.Args) > 1 {
		dir, err = filepath.Abs(os.Args[1])
	} else {
		dir, err = os.Getwd()
	}
	if err != nil {
		panic(err)
	}

	gs := gostatic.New(dir, gostatic.DefaultConfig())

	gs.RenderAll()
}
