# Gostatic Documentation

Gostatic is a simple static page generator supporting pages, articles and projects as content data
types.

## Content data

The page content data is grouped by type into the folders `_articles`, `_pages` and `_projects`.

All data is written as simple Markdown files, using a header started and ended with '---' to define
the required meta-data for the page. The Markdown content below the header is rendered using
[goldmark](https://github.com/yuin/goldmark). You can find the Markdown render code in
`gostatic/markdown.go`.

### Pages

Pages are classic HTML pages using the global site menus and layout. You can use this type e.g. for
your about page.

A page requires the meta-data 'title' and 'subtitle'.

You can find an example page in `test_data/_pages/page.md`.

### Projects

A project is a special HTML page supporting additional meta-data.

A project requires the meta-data 'title', 'subtitle', 'tags' and 'github'. The meta-data 'tags' is
a space separated list of tags. The meta-data 'github' is the Github URL to the project code.

You can find an example page in `test_data/_projects/project.md`.

### Articles

An article is a 'blog post'. This type of content is intended for getting displayed as lists showing
the title and a brief abstract of the content.

An article requires the meta-data 'title', 'category', 'tags' and 'date'. The meta-data 'tags' is
a space separated list of tags. The meta-data 'category' is a kind of primary tag and used to group
articles. The meta-data 'date' is intended as publish date of the article.

You can find an example page in `test_data/_articles/article.md`.

## Menus

The rendered menus are defined as simple text lists and placed in the folder `_structure`. Menus
support on hierarchy level and each line in the text file represents one menu entry.

A line (menu entry) is structured as: [one or more spaces]<menu entry name> -- <menu entry URL>

If a line starts with one or more spaces, it is handled as sub-menu entry of the last menu entry
before.

You can find an example menu in `test_data/_structure/top`.

The name of the file is used to identify the menu data in the templates.

## Design

Gostatic is using the go HTML render feature. You can find more details about the template syntax
at [Go html/template](https://pkg.go.dev/html/template).

The templates for the rendered page are places into the `_templates` folder. All shared templates
are placed into `_templates/snippets`.

Article detail pages are rendered using `_templates/article.tmpl` and all `_templates/snippets`.

Article lists are rendered using `_templates/articles.tmpl` and all `_templates/snippets`.

Article category lists are rendered using `_templates/category.tmpl` and all `_templates/snippets`.

The home page is rendered using `_templates/index.tmpl` and all `_templates/snippets`.

The other pages are rendered using `_templates/page.tmpl` and all `_templates/snippets`.

The special page 'pageMap' is rendered using `_templates/pageMap.tmpl` and all `_templates/snippets`.

The pages list is rendered using `_templates/pages.tmpl` and all `_templates/snippets`.

The project pages are rendered using `_templates/project.tmpl` and all `_templates/snippets`.

The projects list is rendered using `_templates/projects.tmpl` and all `_templates/snippets`.

The tag list is rendered using `_templates/tag.tmpl` and all `_templates/snippets`.

### Meta-data

Additional meta-data form the data file headers can be used in the templates with
`{{index .Meta "<lower case name>"}}`.
