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

	if gs.templates == nil {
		log.Println("TEST: templates is nil")
		t.Fail()
	}

	if gs.structure == nil {
		log.Println("TEST: structure is nil")
		t.Fail()
	}

	// TODO: extend test!
}
