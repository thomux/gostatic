package gostatic

import (
	"html/template"
)

type Gostatic struct {
	root      string
	config    GostaticConfig
	templates map[string]*template.Template
	structure map[string][]Menu
}

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
