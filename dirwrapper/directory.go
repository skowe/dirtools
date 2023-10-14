package dirwrapper

import (
	"os"
)

// Wrapps the *os.File object and provides the basic utilities by which one can check for changes in the contents.
type DirectoryWrapper struct {
	Dir      *os.File
	Contents map[string]string
}
