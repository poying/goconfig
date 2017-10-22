package loader

import (
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// NewYAMLLoader creates a YAML loader
func NewYAMLLoader() *YAMLLoader {
	return &YAMLLoader{}
}

// YAMLLoader represents YAML loader
type YAMLLoader struct{}

// MatchFile implements Loader interface
func (loader *YAMLLoader) MatchFile(file string) bool {
	ext := filepath.Ext(file)
	ext = strings.ToLower(ext)
	return ext == ".yaml" || ext == ".yml"
}

// Unmarshal implements Loader interface
func (loader *YAMLLoader) Unmarshal(b []byte, dest interface{}) error {
	return yaml.Unmarshal(b, dest)
}
