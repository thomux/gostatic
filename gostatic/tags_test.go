package gostatic

import (
	"log"
	"testing"
)

func TestParseTags(t *testing.T) {
	gs := New("../test_data/", DefaultConfig())

	if len(gs.tags) != 6 {
		log.Println("TEST: not 6 tags")
		t.Fail()
	}
}
