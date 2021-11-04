package gostatic

import (
	"bytes"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

// Markdown groups all data about a page or article.
type Markdown struct {
	// Meta contains the meta-data values, e.g. title.
	Meta map[string]interface{}
	// Content contains the Markdown page content.
	Content string
	// Html contains the page contant rendered as HTML.
	Html string
}

// mdToHtml renders the given Markdown string as HTML.
func mdToHtml(content string) string {
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.Table,
			extension.Strikethrough,
			extension.DefinitionList,
			extension.TaskList,
			extension.Footnote,
			extension.Typographer,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	var buf bytes.Buffer
	if err := md.Convert([]byte(content), &buf); err != nil {
		panic(err)
	}

	return buf.String()
}

// parseMarkdown reads the given file as Markdown structure.
// It also reads the metadata header, if available.
func parseMarkdown(path string) Markdown {
	log.Println("Parsing", path)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	meta := make(map[string]interface{})

	offset := 0
	if strings.HasPrefix(lines[0], "---") {
		for i, line := range lines[1:] {
			if strings.HasPrefix(line, "---") {
				offset = i + 2
				break
			}

			parts := strings.SplitN(line, ":", 2)
			if len(parts) > 1 {
				meta[strings.ToLower(strings.TrimSpace(parts[0]))] = strings.TrimSpace(parts[1])
			} else {
				meta[strings.ToLower(strings.TrimSpace(parts[0]))] = ""
			}
		}
	}

	content := strings.Join(lines[offset:], "\n")

	html := content
	file := filepath.Base(path)
	if strings.HasSuffix(file, ".html") {
		meta["file"] = file[:len(file)-5]
	} else if strings.HasSuffix(file, ".md") {
		html = mdToHtml(content)
		meta["file"] = file[:len(file)-3]
	} else {
		panic("extension not supported" + path)
	}

	return Markdown{
		Meta:    meta,
		Content: content,
		Html:    html,
	}
}
