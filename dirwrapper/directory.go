package dirwrapper

import (
	"os"
	"sort"

	"golang.org/x/exp/slices"
)

// Wrapps the *os.File object and provides the basic utilities by which one can check for changes in the contents.
type Directory struct {
	Dir      string
	Contents []string
}

// Returns map with difference between directory reads and an erorr value
// Run the update method after this if differences were detected
// If error exists then the map is empty never nil
// If there are no detectable changes the map is empty and error value is nil
// TODO: Implement a way for Read to go through subdirectories right now it just skips them
func (d *Directory) CheckForUpdate() ([]string, error) {
	//res :=
	f, err := os.Open(d.Dir)
	if err != nil {
		return []string{}, err
	}
	defer f.Close()

	tmp, err := f.ReadDir(0)
	if err != nil {
		return []string{}, err
	}

	curCon := func() []string {
		x := make([]string, 0)
		for _, v := range tmp {
			x = append(x, v.Name())
		}

		sort.Slice(x, func(i, j int) bool {
			return x[i] < x[j]
		})

		return x
	}()
	res := make([]string, 0)
	if len(curCon) > len(d.Contents) {
		for _, file := range curCon {
			if slices.Contains[[]string, string](d.Contents, file) {
				res = append(res, file)
			}
		}
	}
	return res, nil
}
