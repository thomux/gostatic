package gostatic

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// parseProjects reads all project files and initializes the data structures.
func (gs *Gostatic) parseProjects() {
	projectsPath := filepath.Join(gs.root, gs.config.ProjectsPath)

	log.Println("Projects:", projectsPath)

	projectsFiles, err := os.ReadDir(projectsPath)
	if err != nil {
		panic(err)
	}

	for _, entry := range projectsFiles {
		if entry.IsDir() {
			continue
		}

		path := filepath.Join(projectsPath, entry.Name())

		project := parseMarkdown(path)

		tags := strings.Split(project.Meta["tags"].(string), " ")
		project.Meta["tags"] = tags

		project.Meta["url"] = "/" + projectPath(project.Meta["file"].(string))

		gs.projects[entry.Name()] = project

		log.Println("Project:", entry.Name())
	}
}

// RenderProjects renders all project pages.
func (gs *Gostatic) RenderProjects() {
	for file, project := range gs.projects {
		log.Println("Render project", project.Meta["title"], "from", file)
		output := projectPath(project.Meta["file"].(string))
		gs.Render("project.tmpl", output, project.Html, project.Meta)
	}
}

// projectPath returns the path to the rendered project file.
func projectPath(file string) string {
	return filepath.Join("project", file+".html")
}

// articlesByTag provides a map with all articles grouped by tag.
func (gs *Gostatic) projectsByTag() map[string][]Markdown {
	tags := make(map[string][]Markdown)
	for _, v := range gs.projects {
		for _, w := range v.Meta["tags"].([]string) {
			list, ok := tags[w]
			if !ok {
				list = make([]Markdown, 0)
			}
			list = append(list, v)
			tags[w] = list
		}
	}
	return tags
}
