package gostatic

import (
	"log"
	"strings"
	"testing"
	"time"
)

func TestParseArticles(t *testing.T) {
	gs := New("../test_data/", DefaultConfig())
	gs.parseArticles()

	if len(gs.articles) != 4 {
		log.Println("TEST: not 4 articles")
		t.Fail()
	}

	article := gs.articles["article.md"]
	if article.Meta["title"].(string) != "A first article" {
		log.Println("TEST: title wrong")
		t.Fail()
	}
	if article.Meta["category"].(string) != "Server" {
		log.Println("TEST: category wrong")
		t.Fail()
	}
	if len(article.Meta["tags"].([]string)) != 2 {
		log.Println("TEST: tag count wrong")
		t.Fail()
	}
	if len(article.Content) == 0 {
		log.Println("TEST: no content")
		t.Fail()
	}
	if len(article.Html) == 0 {
		log.Println("TEST: no html")
		t.Fail()
	}
}

func TestCategoryPath(t *testing.T) {
	path := categoryPath("Hello")
	if !strings.Contains(path, "category") {
		log.Println("TEST: category not in path")
		t.Fail()
	}
	if !strings.Contains(path, "hello") {
		log.Println("TEST: hello not in path")
		t.Fail()
	}

	path = categoryPath("Hello World")
	if !strings.Contains(path, "hello_world") {
		log.Println("TEST: hello_world not in path")
		t.Fail()
	}
}

func TestTagPath(t *testing.T) {
	path := tagPath("Hello")
	if !strings.Contains(path, "tag") {
		log.Println("TEST: tag not in path")
		t.Fail()
	}
	if !strings.Contains(path, "hello") {
		log.Println("TEST: hello not in path")
		t.Fail()
	}
}

func TestArticlePath(t *testing.T) {
	path := articlePath("Hello")
	if !strings.Contains(path, "article") {
		log.Println("TEST: article not in path")
		t.Fail()
	}
	if !strings.Contains(path, "Hello") {
		log.Println("TEST: Hello not in path")
		t.Fail()
	}
}

func TestArticlesSortedByDate(t *testing.T) {
	gs := New("../test_data/", DefaultConfig())
	gs.parseArticles()

	articles := gs.articlesSortedByDate()
	a := articles[0]
	for _, b := range articles[1:] {
		ta := a.Meta["time"].(time.Time)
		tb := b.Meta["time"].(time.Time)
		if ta.Before(tb) {
			log.Println("TEST: wrong order of", a, "and", b)
			t.Fail()
		}
		a = b
	}
}

func TestLatestArticles(t *testing.T) {
	gs := New("../test_data/", DefaultConfig())
	gs.parseArticles()

	articles := gs.latestArticles()

	if len(articles) > 5 {
		log.Println("TEST: too many articles")
		t.Fail()
	}

	a := articles[0]
	for _, b := range articles[1:] {
		ta := a.Meta["time"].(time.Time)
		tb := b.Meta["time"].(time.Time)
		if ta.Before(tb) {
			log.Println("TEST: wrong order of", a, "and", b)
			t.Fail()
		}
		a = b
	}
}

func TestArticlesByCategory(t *testing.T) {
	gs := New("../test_data/", DefaultConfig())
	gs.parseArticles()

	for category, articles := range gs.articlesByCategory() {
		for _, article := range articles {
			if article.Meta["category"].(string) != category {
				log.Println("TEST: category wrong", category, article)
				t.Fail()
			}
		}
	}
}

func TestArticlesByTag(t *testing.T) {
	gs := New("../test_data/", DefaultConfig())
	gs.parseArticles()

	for tag, articles := range gs.articlesByTag() {
		for _, article := range articles {
			found := false
			for _, t := range article.Meta["tags"].([]string) {
				if t == tag {
					found = true
					break
				}
			}
			if !found {
				log.Println("TEST: tag not found", tag, article)
				t.Fail()
			}
		}
	}
}
