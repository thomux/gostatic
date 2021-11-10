package gostatic

import (
	"testing"
)

func TestNewMenu(t *testing.T) {
	m := newMenu()
	if m.Current ||
		m.Dropdown ||
		m.Separator ||
		m.Name != "" ||
		m.Url != "" ||
		m.Childs == nil ||
		len(m.Childs) > 0 {
		t.Fail()
	}
}

func TestParseMenu(t *testing.T) {
	m := parseMenu("Hallo -- https:\\www.google.de")
	if m.Name != "Hallo" {
		t.Fail()
	}
	if m.Url != "https:\\www.google.de" {
		t.Fail()
	}

	m = parseMenu("  Hallo -- https:\\www.google.de")
	if m.Name != "Hallo" {
		t.Fail()
	}
	if m.Url != "https:\\www.google.de" {
		t.Fail()
	}

	m = parseMenu("  Hallo")
	if m.Name != "Hallo" {
		t.Fail()
	}
	if m.Url != "" {
		t.Fail()
	}

	m = parseMenu("---")
	if !m.Separator {
		t.Fail()
	}

	m = parseMenu("  ---")
	if !m.Separator {
		t.Fail()
	}
}

func TestParseStructure(t *testing.T) {
	menu := parseStructure("../test_data/_structure/top")
	if len(menu) != 4 {
		t.Fail()
	}
	if menu[0].Name != "Github" {
		t.Fail()
	}
	if menu[0].Url != "https://github.com/thomux" {
		t.Fail()
	}
	if len(menu[1].Childs) != 3 {
		t.Fail()
	}
	if menu[3].Childs[1].Separator != true {
		t.Fail()
	}
}
