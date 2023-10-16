package dirwrapper

import (
	"os"
	"sort"
)

// Open the directory at specified path
// Return a pointer to Directory and error
// If error is not nil Directory points to an empty object
// TODO: Implement a way for Open to go through subdirectories right now it just skips them
func Open(directoryPath string) (*Directory, error) {

	res := &Directory{
		Dir:      "",
		Contents: make([]string, 0),
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

	contents := make([]string, 0)
	for _, v := range dirEntries {
		if err != nil && !v.IsDir() {
			return res, &HashingError{}
		} else if !v.IsDir() {
			contents = append(contents, v.Name())
		}
	}
	sort.Slice(contents, func(i, j int) bool {
		return contents[i] < contents[j]
	})
	res.Dir = directoryPath
	res.Contents = contents
	return res, nil
}

func Make(d *Directory) error {

	return os.MkdirAll(d.Dir, 0766)
}
