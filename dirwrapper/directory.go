package dirwrapper

import (
	"os"
	"sort"

	"golang.org/x/exp/slices"
)

//TODO: WRITE THE GOD DAMN DOCS
type Directory struct {
	Dir      string
	Contents []string
}

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
	d.Contents = res
	return res, nil
}
