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
}

// New creates a new gostatic struct with default values.
func New(root string, config GostaticConfig) *Gostatic {
	gs := new(Gostatic)
	gs.root = root
	gs.templates = make(map[string]*template.Template)
	gs.structure = make(map[string][]Menu)
	gs.config = config

	gs.parseTemplates()
	gs.readStructure()

	return gs
}
