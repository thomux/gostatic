package gostatic

import "testing"

func TestParseTemplates(t *testing.T) {
	config := DefaultConfig()
	gs := New("../test_data", config)
	gs.parseTemplates()
	_, ok := gs.templates["index.tmpl"]
	if !ok {
		t.Fail()
	}
	_, ok = gs.templates["page.tmpl"]
	if !ok {
		t.Fail()
	}
}
