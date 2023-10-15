package dirwrapper

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Open the directory at specified path
// Return a pointer to DirectoryWrapper and error
// If error is not nil DirectoryWrapper points to an empty object
func Open(directoryPath string) (*DirectoryWrapper, error) {

	res := &DirectoryWrapper{
		Dir:      "",
		Contents: make(map[string]string),
	}
	f, err := os.Open(directoryPath)
	if err != nil {
		return res, err
	}
	defer f.Close()

	stat, err := f.Stat()

	if err != nil {
		return res, err
	}

	if !stat.IsDir() {
		return res, &FileNotDirError{}
	}

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
	res.Dir = directoryPath
	res.Contents = contents
	return res, nil
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

func Make(d *DirectoryWrapper) error {

	return os.MkdirAll(d.Dir, 0766)
}
