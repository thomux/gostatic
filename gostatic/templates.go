package gostatic

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
)

// parseSnippets loads all template snippets.
func parseSnippets(path string) *template.Template {
	snippets, err := template.ParseGlob(path)
	if err != nil {
		panic(err)
	}
	return snippets
}

// parseTemplates load all templates.
func (gs *Gostatic) parseTemplates() {
	templatesPath := filepath.Join(gs.root, gs.config.TemplatePath)
	snippetsPath := filepath.Join(gs.root, gs.config.SnippetsPath, "*")

	log.Println("Snippets:", snippetsPath)
	log.Println("Templates:", templatesPath)

	templateFiles, err := os.ReadDir(templatesPath)
	if err != nil {
		panic(err)
	}

	for _, entry := range templateFiles {
		if entry.IsDir() {
			continue
		}

		path := filepath.Join(templatesPath, entry.Name())

		snippets := parseSnippets(snippetsPath)

		t, err := snippets.ParseFiles(path)
		if err != nil {
			panic(err)
		}
		gs.templates[entry.Name()] = t

		log.Println("Template:", entry.Name())
	}
}
