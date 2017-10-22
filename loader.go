package goconfig

// Loader represents loader interface
type Loader interface {
	MatchFile(filepath string) bool
	Unmarshal([]byte, interface{}) error
}
