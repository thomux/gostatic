package gostatic

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

type Data struct {
	Title     string
	Structure map[string][]Menu
}

func prepareDir(path string) {
	parent := filepath.Dir(path)
	log.Println("Create dir", parent)
	err := os.MkdirAll(parent, 0770)
	if err != nil {
		panic(err)
	}
}

func fileWriter(path string) (*os.File, *bufio.Writer) {
	// Remove file if exists
	_ = os.Remove(path)
	// Create new file for writing
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return f, bufio.NewWriter(f)
}

func (gs *Gostatic) Render(template string, file string, data interface{}) {
	t := gs.templates[template]

	path := filepath.Join(gs.root, gs.config.Output, file)

	prepareDir(path)
	f, w := fileWriter(path)

	// TODO: select current

	err := t.ExecuteTemplate(w, template, Data{
		Title:     "Thomux",
		Structure: gs.structure,
	})
	if err != nil {
		panic(err)
	}
	err = w.Flush()
	if err != nil {
		panic(err)
	}
	f.Close()

	log.Println("Rendered", template, "to", path)
}
