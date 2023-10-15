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

// Returns map with difference between directory reads and an erorr value
// Run the update method after this if differences were detected
// If error exists then the map is empty never nil
// If there are no detectable changes the map is empty and error value is nil
func (d *DirectoryWrapper) Read() (map[string]string, error) {
	res := make(map[string]string)
	f, err := os.Open(d.Dir)
	if err != nil {
		return res, err
	}
	defer f.Close()

	dirEntries, err := f.ReadDir(0)
	if err != nil {
		return res, err
	}

	for _, v := range dirEntries {
		h, err := hash(filepath.Join(f.Name(), v.Name()))
		if err != nil && !v.IsDir() {
			return res, &HashingError{}
		} else if !v.IsDir() {
			res[v.Name()] = h
		}
	}

	return res, nil
}
