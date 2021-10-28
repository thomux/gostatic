package gostatic

import "testing"

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()
	if config.TemplatePath != "_templates" {
		t.Fail()
	}
	if config.SnippetsPath != "_templates/snippets" {
		t.Fail()
	}
	if config.StructurePath != "_structure" {
		t.Fail()
	}
	if config.Output != "" {
		t.Fail()
	}
}
