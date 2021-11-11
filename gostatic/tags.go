package gostatic

import (
	"log"
)

// parseTags prepares the tag data structures.
// It must be called after parseArticles and parseProjects to have the required
// data available.
func (gs *Gostatic) parseTags() {
	gs.tags = gs.tagUrls()
}

// RenderTagPages renders all tag pages.
func (gs *Gostatic) RenderTagPages() {
	projectsByTag := gs.projectsByTag()
	articlesByTag := gs.articlesByTag()

	for _, link := range gs.tags {
		tag := link.Name
		log.Println("Render tag page", tag)
		output := tagPath(tag)
		meta := make(map[string]interface{})
		meta["title"] = tag

		if articles, ok := articlesByTag[tag]; ok {
			meta["articles"] = articles
		} else {
			meta["articles"] = make([]Markdown, 0)
		}

		if projects, ok := projectsByTag[tag]; ok {
			meta["projects"] = projects
		} else {
			meta["projects"] = make([]Markdown, 0)
		}

		gs.Render("tag.tmpl", output, "", meta)
	}
}
