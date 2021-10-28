package main

import (
	"os"
	"path/filepath"

	"github.com/thomux/gostatic/gostatic"
)

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

	gs.Render("index.tmpl", "index.html", nil)
	gs.Render("page.tmpl", "pages/page.html", nil)
}
