package config

type Config struct {
	Project       string `yaml:"-"`
	Root          string `yaml:"root"`
	AdrTemplate   string `yaml:"adrTemplate"`
	IndexTemplate string `yaml:"indexTemplate"`
}
