package gostatic

import (
	"log"
	"strings"
	"testing"
)

func TestParseProjects(t *testing.T) {
	gs := New("../test_data/", DefaultConfig())

	if len(gs.projects) != 1 {
		log.Println("TEST: page count error")
		t.Fail()
	}

	page := gs.projects["project.md"]
	if page.Meta["title"].(string) != "A Project" {
		log.Println("TEST: title wrong")
		t.Fail()
	}
	if page.Meta["subtitle"].(string) != "This is a project!" {
		log.Println("TEST: subtitle wrong")
		t.Fail()
	}
	if len(page.Content) == 0 {
		log.Println("TEST: no content")
		t.Fail()
	}
	if len(page.Html) == 0 {
		log.Println("TEST: no html")
		t.Fail()
	}
}

func TestProjectPath(t *testing.T) {
	path := projectPath("Hello")
	if !strings.Contains(path, "project") {
		log.Println("TEST: project not in path")
		t.Fail()
	}
	if !strings.Contains(path, "Hello") {
		log.Println("TEST: Hello not in path")
		t.Fail()
	}
}

func TestProjectsByTag(t *testing.T) {
	gs := New("../test_data/", DefaultConfig())
	gs.parseProjects()

	for tag, projects := range gs.projectsByTag() {
		for _, project := range projects {
			found := false
			for _, t := range project.Meta["tags"].([]string) {
				if t == tag {
					found = true
					break
				}
			}
			if !found {
				log.Println("TEST: tag not found", tag, project)
				t.Fail()
			}
		}
	}
}
