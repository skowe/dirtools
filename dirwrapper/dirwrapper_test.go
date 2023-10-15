package dirwrapper

import (
	"testing"
)

func TestMake(t *testing.T) {
	myDir1 := &DirectoryWrapper{
		Dir:      "/home/skowe/Projects/dirtools/dirwrapper/dosntexist/dir1",
		Contents: make(map[string]string),
	}

	myDir2 := &DirectoryWrapper{
		Dir:      "/home/skowe/Projects/dirtools/dirwrapper/exists",
		Contents: make(map[string]string),
	}

	myDir3 := &DirectoryWrapper{
		Dir:      "/home/skowe/Projects/dirtools/dirwrapper/exists/subdir1",
		Contents: make(map[string]string),
	}

	err := Make(myDir1)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	err = Make(myDir2)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	err = Make(myDir3)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
