package gostatic

type GostaticConfig struct {
	SnippetsPath  string
	TemplatePath  string
	StructurePath string
	Output        string
}

func DefaultConfig() GostaticConfig {
	return GostaticConfig{
		SnippetsPath:  "_templates/snippets",
		TemplatePath:  "_templates",
		StructurePath: "_structure",
		Output:        "",
	}
}
