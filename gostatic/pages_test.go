package gostatic

import (
	"log"
	"strings"
	"testing"
)

func TestParsePages(t *testing.T) {
	gs := New("../test_data/", DefaultConfig())

	if len(gs.pages) != 1 {
		log.Println("TEST: page count error")
		t.Fail()
	}

	page := gs.pages["page.md"]
	if page.Meta["title"].(string) != "A Page" {
		log.Println("TEST: title wrong")
		t.Fail()
	}
	if page.Meta["subtitle"].(string) != "This is a page!" {
		log.Println("TEST: category wrong")
		t.Fail()
	}
	if len(page.Content) == 0 {
		log.Println("TEST: no content")
		t.Fail()
	}
	if len(page.Html) == 0 {
		log.Println("TEST: no html")
		t.Fail()
	}
}

func TestPagePath(t *testing.T) {
	path := pagePath("Hello")
	if !strings.Contains(path, "page") {
		log.Println("TEST: page not in path")
		t.Fail()
	}
	if !strings.Contains(path, "Hello") {
		log.Println("TEST: Hello not in path")
		t.Fail()
	}
}

func TestPageUrls(t *testing.T) {
	gs := New("../test_data/", DefaultConfig())

	links := gs.pageUrls()

	if len(links) != 1 {
		log.Println("TEST: wrong link count")
		t.Fail()
	}

	if links[0].Name != "A Page" {
		log.Println("TEST: wrong page name")
		t.Fail()
	}

	if links[0].Url != "/page/page.html" {
		log.Println("TEST: wrong page URL")
		t.Fail()
	}
}

func TestProjectUrls(t *testing.T) {
	gs := New("../test_data/", DefaultConfig())

	links := gs.projectUrls()

	if len(links) != 1 {
		log.Println("TEST: wrong link count")
		t.Fail()
	}

	if links[0].Name != "A Project" {
		log.Println("TEST: wrong project name")
		t.Fail()
	}

	if links[0].Url != "/project/project.html" {
		log.Println("TEST: wrong project URL")
		t.Fail()
	}
}

func TestArticleUrls(t *testing.T) {
	gs := New("../test_data/", DefaultConfig())

	links := gs.articleUrls()

	if len(links) != 4 {
		log.Println("TEST: wrong link count")
		t.Fail()
	}

	if !strings.Contains(links[0].Name, " article") {
		log.Println("TEST: wrong article name")
		t.Fail()
	}

	if !strings.Contains(links[0].Url, "/article/article") {
		log.Println("TEST: wrong article URL")
		t.Fail()
	}
}

func TestCategoryUrls(t *testing.T) {
	gs := New("../test_data/", DefaultConfig())

	links := gs.categoryUrls()

	if len(links) != 3 {
		log.Println("TEST: wrong link count")
		t.Fail()
	}

	if !strings.Contains(links[0].Url, "/category/") {
		log.Println("TEST: wrong category URL")
		t.Fail()
	}
}

func TestTagUrls(t *testing.T) {
	gs := New("../test_data/", DefaultConfig())

	links := gs.tagUrls()

	if len(links) != 6 {
		log.Println("TEST: wrong link count")
		t.Fail()
	}

	if !strings.Contains(links[0].Url, "/tag/") {
		log.Println("TEST: wrong tag URL")
		t.Fail()
	}
}
