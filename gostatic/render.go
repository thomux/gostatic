package gostatic

import (
	"bufio"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Data is the structure used by the templates.
type Data struct {
	// Title of the page.
	Title string
	// Content is the page content.
	Content template.HTML
	// Menu data structures.
	Structure map[string][]Menu
	// Meta contains further meta-data.
	Meta map[string]interface{}
	// Categories is a link list for all categories.
	Categories []Link
	// Tags is a link list for all tags.
	Tags []Link
}

// prepareDir recursive creates a directory path for the given file.
// It panics on error.
func prepareDir(path string) {
	parent := filepath.Dir(path)
	log.Println("Create dir", parent)
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
func (gs *Gostatic) Render(template string, file string, content string, meta map[string]interface{}) {
	t := gs.templates[template]

	path := filepath.Join(gs.root, gs.config.Output, file)

	prepareDir(path)
	f, w := fileWriter(path)

	gs.selectCurrentPage(meta["title"].(string))

	err := t.ExecuteTemplate(w, template, Data{
		Title:      meta["title"].(string),
		Content:    safe(content),
		Structure:  gs.structure,
		Meta:       meta,
		Categories: gs.categories,
		Tags:       gs.tags,
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

// safe wraps a string into a template.HTML to mark it as safe content.
func safe(content string) template.HTML {
	return template.HTML(content)
}

// RenderAll renders all data, i.e. the index page, all other pages,
// all articles, all projects, and all special pages.
func (gs *Gostatic) RenderAll() {
	log.Println("Render index page ...")
	gs.RenderIndex()
	log.Println("Render other page ...")
	gs.RenderPages()
	log.Println("Render articles ...")
	gs.RenderArticles()
	log.Println("Render projects ...")
	gs.RenderProjects()
	log.Println("Render special pages ...")
	gs.RenderSpecial()
	log.Println("Render tag pages ...")
	gs.RenderTagPages()
}
