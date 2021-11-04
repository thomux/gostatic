/*
gostatic implements all logic of the static site generator.
*/
package gostatic

import (
	"html/template"
)

// Gostatic groups all runtime information.
type Gostatic struct {
	// root contains the path of the data folder.
	root string
	// config groups default configurations.
	config GostaticConfig
	// templates is a map of all found templates.
	templates map[string]*template.Template
	// structure is a map of all found menu data.
	structure map[string][]Menu
	// pages is a map of all found page data.
	pages map[string]Markdown
	// articles is a map of all found article data.
	articles map[string]Markdown
	// articles is a map of all found project data.
	projects map[string]Markdown
	// categories is a list of all known categories.
	categories []Link
	// tags is a list of all known tags.
	tags []Link
}

// New creates a new gostatic struct with default values.
func New(root string, config GostaticConfig) *Gostatic {
	gs := new(Gostatic)
	gs.root = root
	gs.templates = make(map[string]*template.Template)
	gs.structure = make(map[string][]Menu)
	gs.pages = make(map[string]Markdown)
	gs.articles = make(map[string]Markdown)
	gs.projects = make(map[string]Markdown)
	gs.tags = make([]Link, 0)
	gs.categories = make([]Link, 0)
	gs.config = config

	gs.parseTemplates()
	gs.readStructure()
	gs.parsePages()
	gs.parseArticles()
	gs.parseProjects()

	return gs
}
