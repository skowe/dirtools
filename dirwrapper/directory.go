package dirwrapper

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Wrapps the *os.File object and provides the basic utilities by which one can check for changes in the contents.
type DirectoryWrapper struct {
	Dir      *os.File
	Contents map[string]string
}

// Open the directory at specified path
// Return a pointer to DirectoryWrapper and error
// If error is not nil DirectoryWrapper points to an empty object
func Open(directoryPath string) (*DirectoryWrapper, error) {
	f, err := os.Open(directoryPath)

	if err != nil {
		return &DirectoryWrapper{
			Dir:      nil,
			Contents: make(map[string]string),
		}, err
	}
	defer f.Close()

	stat, err := f.Stat()

	if err != nil {
		return &DirectoryWrapper{
			Dir:      nil,
			Contents: make(map[string]string),
		}, err
	}

	if !stat.IsDir() {
		return &DirectoryWrapper{
			Dir:      nil,
			Contents: make(map[string]string),
		}, &FileNotDirError{}
	}

	dirEntries, err := f.ReadDir(0)
	if err != nil {
		return &DirectoryWrapper{
			Dir:      nil,
			Contents: make(map[string]string),
		}, err
	}

	contents := make(map[string]string)
	for _, v := range dirEntries {
		h, err := hash(filepath.Join(f.Name(), v.Name()))
		if err != nil {
			return &DirectoryWrapper{
				Dir:      nil,
				Contents: make(map[string]string),
			}, &HashingError{}
		}
		contents[v.Name()] = h
	}

	return &DirectoryWrapper{
		Dir:      f,
		Contents: contents,
	}, nil
}

func hash(pathToFile string) (string, error) {
	f, err := os.Open(pathToFile)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()

	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
