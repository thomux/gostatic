package gostatic

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// parseArticles reads all article files and initializes the data structures.
func (gs *Gostatic) parseArticles() {
	pagesPath := filepath.Join(gs.root, gs.config.ArticlesPath)

	log.Println("Articles:", pagesPath)

	pagesFiles, err := os.ReadDir(pagesPath)
	if err != nil {
		panic(err)
	}

	for _, entry := range pagesFiles {
		if entry.IsDir() {
			continue
		}

		path := filepath.Join(pagesPath, entry.Name())

		article := parseMarkdown(path)

		tags := strings.Split(article.Meta["tags"].(string), " ")
		article.Meta["tags"] = tags

		idx := strings.Index(article.Html, "</p>")
		if idx == -1 {
			article.Meta["summary"] = template.HTML(article.Html)
			article.Content = ""
		} else {
			article.Meta["summary"] = template.HTML(article.Html[:idx+4])
			article.Content = article.Html[idx+4:]
		}

		t, err := time.Parse("2006-01-02", article.Meta["date"].(string))
		if err != nil {
			panic(err)
		}
		article.Meta["time"] = t

		gs.articles[entry.Name()] = article

		article.Meta["url"] = "/" + articlePath(article.Meta["file"].(string))

		log.Println("Article:", entry.Name())
	}

	gs.categories = gs.categoryUrls()
	gs.tags = gs.tagUrls()
}

// RenderArticles renders all articles.
func (gs *Gostatic) RenderArticles() {
	for file, article := range gs.articles {
		log.Println("Render article", article.Meta["title"], "from", file)
		output := article.Meta["url"].(string)[1:]
		gs.Render("article.tmpl", output, article.Html, article.Meta)
	}

	for category, articles := range gs.articlesByCategory() {
		log.Println("Render category page", category)
		output := categoryPath(category)
		meta := make(map[string]interface{})
		meta["title"] = category
		meta["articles"] = articles
		gs.Render("category.tmpl", output, "", meta)
	}

	for tag, articles := range gs.articlesByTags() {
		log.Println("Render tag page", tag)
		output := tagPath(tag)
		meta := make(map[string]interface{})
		meta["title"] = tag
		meta["articles"] = articles
		gs.Render("category.tmpl", output, "", meta)
	}
}

// categoryPath returns the path to the file listing the articles
// of the given category.
func categoryPath(category string) string {
	return filepath.Join(
		"category",
		strings.ToLower(strings.ReplaceAll(category, " ", "_"))+".html")
}

// categoryPath returns the path to the file listing the articles
// with the given tag.
func tagPath(tag string) string {
	return filepath.Join("tag", strings.ToLower(tag)+".html")
}

// articlePath returns the path to the rendered article file.
func articlePath(file string) string {
	return filepath.Join("article", file+".html")
}

// Render the index page, containing the latest articles.
func (gs *Gostatic) RenderIndex() {
	meta := make(map[string]interface{})
	meta["title"] = "Home"
	meta["articles"] = gs.latestArticles()
	gs.Render("index.tmpl", "index.html", "", meta)
}

// articlesSortedByDate generates a list of all articles sorted by date.
// The newest article is the first (index 0).
func (gs *Gostatic) articlesSortedByDate() []Markdown {
	articles := make([]Markdown, len(gs.articles))
	i := 0
	for _, a := range gs.articles {
		articles[i] = a
		i++
	}

	sort.Slice(articles, func(i, j int) bool {
		a := articles[i].Meta["time"].(time.Time)
		b := articles[j].Meta["time"].(time.Time)
		return a.After(b)
	})

	return articles
}

// latestArticles provides a list containing the latest articles.
func (gs *Gostatic) latestArticles() []Markdown {
	articles := gs.articlesSortedByDate()
	if len(articles) > 5 {
		return articles[:5]
	} else {
		return articles
	}
}

// articlesByCategory provides a map with all articles grouped by category.
func (gs *Gostatic) articlesByCategory() map[string][]Markdown {
	categories := make(map[string][]Markdown)
	for _, v := range gs.articles {
		category := v.Meta["category"].(string)
		list, ok := categories[category]
		if !ok {
			list = make([]Markdown, 0)
		}
		list = append(list, v)
		categories[category] = list
	}
	return categories
}

// articlesByTags provides a map with all articles grouped by tag.
func (gs *Gostatic) articlesByTags() map[string][]Markdown {
	tags := make(map[string][]Markdown)
	for _, v := range gs.articles {
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
