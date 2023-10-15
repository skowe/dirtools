package dirwrapper

import (
	"os"
	"path/filepath"
)

// Wrapps the *os.File object and provides the basic utilities by which one can check for changes in the contents.
type DirectoryWrapper struct {
	Dir      string
	Contents map[string]string
}

func (d *DirectoryWrapper) Read() (os.DirEntry, error) {

	f, err := os.Open(d.Dir)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dirEntries, err := f.ReadDir(0)
	if err != nil {
		return res, err
	}

	contents := make(map[string]string)
	for _, v := range dirEntries {
		h, err := hash(filepath.Join(f.Name(), v.Name()))
		if err != nil {
			return res, &HashingError{}
		}
		contents[v.Name()] = h
	}
}
