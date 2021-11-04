package gostatic

// GostaticConfig groups all config values.
type GostaticConfig struct {
	// SnippetsPath is the path to the template snipptes.
	SnippetsPath string
	// TemplatePath is the path to the templates.
	TemplatePath string
	// StructurePath is the path to the menu structures.
	StructurePath string
	// Output the the base folder for the generated pages.
	Output string
}

// DefaultConfig creates a new GostaticConfig initialized with
// the defaut paths for the data folders.
func DefaultConfig() GostaticConfig {
	return GostaticConfig{
		SnippetsPath:  "_templates/snippets",
		TemplatePath:  "_templates",
		StructurePath: "_structure",
		Output:        "",
	}
}
