package gostatic

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPrepareDir(t *testing.T) {
	path := "./hallo/welt/file.text"
	parent := filepath.Dir(path)
	prepareDir(path)
	info, err := os.Stat(parent)
	if err != nil {
		log.Println("TEST:", err)
		t.Fail()
	}

	if !info.IsDir() {
		log.Println("TEST:", parent, "is no dir")
		t.Fail()
	}

	err = os.RemoveAll(parent)
	if err != nil {
		log.Println("TEST: cleanup", err)
		t.Fail()
	}

	err = os.RemoveAll(filepath.Dir(parent))
	if err != nil {
		log.Println("TEST: cleanup", err)
		t.Fail()
	}
}

func TestSelectCurrentPage(t *testing.T) {
	root := "../test_data/"
	config := DefaultConfig()
	gs := New(root, config)

	gs.selectCurrentPage("Server")

	for _, ms := range gs.structure {
		for _, m := range ms {
			if strings.ToLower(m.Name) == "server" {
				if !m.Current {
					log.Println("TEST:", m.Name, "not current")
					t.Fail()
				}
			}
			for _, sm := range m.Childs {
				if strings.ToLower(sm.Name) == "server" {
					if !sm.Current {
						log.Println("TEST:", sm.Name, "not current")
						t.Fail()
					}

					if !m.Current {
						log.Println("TEST: parent not current")
						t.Fail()
					}
				}
			}
		}
	}
}
