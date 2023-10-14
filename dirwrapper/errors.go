package dirwrapper

type FileNotDirError struct {
}
type HashingError struct {
}

func (e *FileNotDirError) Error() string {
	return "file specified is not a directory"
}

func (e *HashingError) Error() string {
	return "failed to hash file contents"
}
