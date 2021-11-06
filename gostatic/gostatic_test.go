package gostatic

import (
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	root := "../test_data/"
	config := DefaultConfig()
	gs := New(root, config)

	if gs.root != root {
		log.Println("TEST: root path wrong")
		t.Fail()
	}

	if gs.config != config {
		log.Println("TEST: config wrong")
		t.Fail()
	}

	if len(gs.templates) != 10 {
		log.Println("TEST: templates missing")
		t.Fail()
	}

	if len(gs.structure) != 2 {
		log.Println("TEST: structure missing")
		t.Fail()
	}

	if len(gs.pages) != 1 {
		log.Println("TEST: page missing")
		t.Fail()
	}

	if len(gs.articles) != 4 {
		log.Println("TEST: articles missing")
		t.Fail()
	}

	if len(gs.projects) != 1 {
		log.Println("TEST: projects missing")
		t.Fail()
	}

	if len(gs.tags) != 6 {
		log.Println("TEST: tags missing")
		t.Fail()
	}

	if len(gs.categories) != 3 {
		log.Println("TEST: categories missing")
		t.Fail()
	}
}
