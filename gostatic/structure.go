package gostatic

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Menu describes one menu entry.
type Menu struct {
	// Separator is true if a menu separator instead of the entry
	// should be rendered.
	Separator bool
	// Dropdown is true if a submenu should be rendered.
	Dropdown bool
	// Current is true if this is the current meny entry.
	Current bool
	// Name is the displayed text of the menu entry.
	Name string
	// Url is the URL of this menu entry.
	Url string
	// Childs is a list of all submenu entries.
	Childs []Menu
}

// newMenu creates a new Menu, initialized with default values.
func newMenu() Menu {
	return Menu{
		Separator: false,
		Dropdown:  false,
		Current:   false,
		Name:      "",
		Url:       "",
		Childs:    make([]Menu, 0),
	}
}

// readStructure reads all menu structure files.
func (gs *Gostatic) readStructure() {
	structurePath := filepath.Join(gs.root, gs.config.StructurePath)

	log.Println("Structure:", structurePath)

	structureFiles, err := os.ReadDir(structurePath)
	if err != nil {
		panic(err)
	}

	for _, entry := range structureFiles {
		if entry.IsDir() {
			continue
		}

		path := filepath.Join(structurePath, entry.Name())

		menus := parseStructure(path)
		gs.structure[entry.Name()] = menus
	}
}

// parseStructure creates the Menu list described in the structure file.
func parseStructure(path string) []Menu {
	log.Println("Parsing", path)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	result := make([]Menu, 0)
	for _, line := range lines {
		if len(line) > 0 && line[0] == '#' {
			log.Println("Skipping comment", line)
			continue
		}
		if len(result) > 0 && len(line) > 0 && line[0] == ' ' {
			last := result[len(result)-1]

			m := parseMenu(line)
			last.Childs = append(last.Childs, m)

			last.Dropdown = true

			result[len(result)-1] = last
		} else {
			result = append(result, parseMenu(line))
		}
	}

	return result
}

// parseMenu creates the Menu described by one line of the structure file.
func parseMenu(line string) Menu {
	m := newMenu()
	parts := strings.Split(line, " -- ")
	if len(parts) > 0 {
		m.Name = strings.TrimSpace(parts[0])
		if m.Name == "---" {
			m.Separator = true
		}
	}
	if len(parts) > 1 {
		m.Url = strings.TrimSpace(parts[1])
	}

	return m
}
