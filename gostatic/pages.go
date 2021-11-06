package gostatic

import (
	"log"
	"os"
	"path/filepath"
)

// Link groups a URL with a name.
type Link struct {
	// Name of the link.
	Name string
	// Url of the link.
	Url string
}

// parsePages reads all article files and initializes the data structures.
func (gs *Gostatic) parsePages() {
	pagesPath := filepath.Join(gs.root, gs.config.PagesPath)

	log.Println("Pages:", pagesPath)

	pagesFiles, err := os.ReadDir(pagesPath)
	if err != nil {
		panic(err)
	}

	for _, entry := range pagesFiles {
		if entry.IsDir() {
			continue
		}

		path := filepath.Join(pagesPath, entry.Name())

		gs.pages[entry.Name()] = parseMarkdown(path)

		log.Println("Page:", entry.Name())
	}
}

// RenderPages renders all pages.
func (gs *Gostatic) RenderPages() {
	for file, page := range gs.pages {
		log.Println("Render page", page.Meta["title"], "from", file)
		output := pagePath(page.Meta["file"].(string))
		gs.Render("page.tmpl", output, page.Html, page.Meta)
	}
}

// pagePath returns the path to the rendered page file.
func pagePath(file string) string {
	return filepath.Join("page", file+".html")
}

// RenderSpecial renders all special pages, i.e. the page map
// and all other lising pages.
func (gs *Gostatic) RenderSpecial() {
	meta := make(map[string]interface{})
	meta["title"] = "Page Map"

	meta["pages"] = gs.pageUrls()
	meta["projects"] = gs.projectUrls()
	meta["articles"] = gs.articleUrls()
	meta["categories"] = gs.categoryUrls()
	meta["tags"] = gs.tagUrls()

	gs.Render("pageMap.tmpl", "pageMap.html", "", meta)

	meta["title"] = "Pages"
	gs.Render("pages.tmpl", "pages.html", "", meta)

	meta["title"] = "Articles"
	meta["articles"] = gs.articlesSortedByDate()
	gs.Render("articles.tmpl", "articles.html", "", meta)

	meta["title"] = "Projects"
	gs.Render("projects.tmpl", "projects.html", "", meta)
}

// pageUrls generates a link list with all rendered pages.
func (gs *Gostatic) pageUrls() []Link {
	links := make([]Link, 0)
	for _, v := range gs.pages {
		links = append(links, Link{
			Name: v.Meta["title"].(string),
			Url:  "/" + pagePath(v.Meta["file"].(string)),
		})
	}
	return links
}

// projectUrls generates a link list with all rendered project pages.
func (gs *Gostatic) projectUrls() []Link {
	links := make([]Link, 0)
	for _, v := range gs.projects {
		links = append(links, Link{
			Name: v.Meta["title"].(string),
			Url:  "/" + projectPath(v.Meta["file"].(string)),
		})
	}
	return links
}

// articleUrls generates a link list with all rendered articles.
func (gs *Gostatic) articleUrls() []Link {
	links := make([]Link, 0)
	for _, v := range gs.articles {
		links = append(links, Link{
			Name: v.Meta["title"].(string),
			Url:  "/" + articlePath(v.Meta["file"].(string)),
		})
	}
	return links
}

// categoryUrls generates a link list with all category list pages.
func (gs *Gostatic) categoryUrls() []Link {
	links := make([]Link, 0)
	for category := range gs.articlesByCategory() {
		links = append(links, Link{
			Name: category,
			Url:  "/" + categoryPath(category),
		})
	}
	return links
}

// tagUrls generates a link list with all tag list pages.
func (gs *Gostatic) tagUrls() []Link {
	links := make([]Link, 0)
	for tag := range gs.articlesByTag() {
		links = append(links, Link{
			Name: tag,
			Url:  "/" + tagPath(tag),
		})
	}
	return links
}
