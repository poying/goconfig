package loader

import (
	"encoding/json"
	"path/filepath"
	"strings"
)

// NewJSONLoader creates a JSON loader
func NewJSONLoader() *JSONLoader {
	return &JSONLoader{}
}

// JSONLoader represents json loader
type JSONLoader struct{}

// MatchFile implements Loader interface
func (loader *JSONLoader) MatchFile(file string) bool {
	ext := filepath.Ext(file)
	ext = strings.ToLower(ext)
	return ext == ".json"
}

// Unmarshal implements Loader interface
func (loader *JSONLoader) Unmarshal(b []byte, dest interface{}) error {
	return json.Unmarshal(b, dest)
}
