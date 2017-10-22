package goconfig

// UnrecognizedFileTypeError represents unrecognized file type error
type UnrecognizedFileTypeError struct {
	Filepath string
}

func (err UnrecognizedFileTypeError) Error() string {
	return "Unrecognized File Type: " + err.Filepath
}

// IOError represents io error
type IOError struct {
	Reason error
}

func (err IOError) Error() string {
	return "IO Error: " + err.Reason.Error()
}
