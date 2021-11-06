package gostatic

import (
	"log"
	"testing"
)

func TestParseMarkdown(t *testing.T) {
	md := parseMarkdown("../test_data/_articles/article.md")
	if len(md.Content) == 0 {
		log.Println("TEST: content wrong")
		t.Fail()
	}

	if len(md.Html) == 0 {
		log.Println("TEST: html wrong")
		t.Fail()
	}

	if md.Meta["title"].(string) != "A first article" {
		log.Println("TEST: title wrong")
		t.Fail()
	}

	if md.Meta["category"].(string) != "Server" {
		log.Println("TEST: category wrong")
		t.Fail()
	}
}

func TestMdToHtml(t *testing.T) {
	data := "Hello\n\nWorld"
	expect := "<p>Hello</p>\n<p>World</p>\n"
	rendered := mdToHtml(data)
	if rendered != expect {
		log.Println("TEST: not equal", rendered)
		t.Fail()
	}
}
