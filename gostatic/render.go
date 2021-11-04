package gostatic

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Data is the structure used by the templates.
type Data struct {
	// Title of the page.
	Title string
	// Menu data structures.
	Structure map[string][]Menu
}

// prepareDir recursive creates a directory path for the given file.
// It panics on error.
func prepareDir(path string) {
	parent := filepath.Dir(path)
	fmt.Println("Create dir", parent)
	err := os.MkdirAll(parent, 0770)
	if err != nil {
		panic(err)
	}
}

// fileWriter creates a new file for buffered writing on the given path.
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

// selectCurrentPage updates the Current flag for the menu structures.
func (gs *Gostatic) selectCurrentPage(title string) {
	t := strings.ToLower(title)

	for key, menus := range gs.structure {
		for i, m := range menus {
			m.Current = strings.ToLower(m.Name) == t

			for j, sm := range m.Childs {
				sm.Current = strings.ToLower(sm.Name) == t
				m.Childs[j] = sm

				if sm.Current {
					m.Current = true
				}
			}

			menus[i] = m
		}
		gs.structure[key] = menus
	}
}

// Render the template to the file using the given template data.
func (gs *Gostatic) Render(template string, file string, title string) {
	t := gs.templates[template]

	path := filepath.Join(gs.root, gs.config.Output, file)

	prepareDir(path)
	f, w := fileWriter(path)

	gs.selectCurrentPage(title)

	err := t.ExecuteTemplate(w, template, Data{
		Title:     title,
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
