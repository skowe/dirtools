package dirwrapper

import "fmt"

type FileNotDirError struct {
}
type HashingError struct {
	fName string
}

func (e *FileNotDirError) Error() string {
	return "file specified is not a directory"
}

func (e *HashingError) Error() string {
	return fmt.Sprintf("failed to hash file %s contents", e.fName)
}
