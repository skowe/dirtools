package dirwrapper

import (
	"testing"
)

func TestMake(t *testing.T) {
	myDir1 := &Directory{
		Dir:      "/home/skowe/Projects/dirtools/dirwrapper/dosntexist/dir1",
		Contents: make([]string, 0),
	}

	myDir2 := &Directory{
		Dir:      "/home/skowe/Projects/dirtools/dirwrapper/exists",
		Contents: make([]string, 0),
	}

	myDir3 := &Directory{
		Dir:      "/home/skowe/Projects/dirtools/dirwrapper/exists/subdir1",
		Contents: make([]string, 0),
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

func TestScan(t *testing.T) {

	myDir, err := Open("/home/skowe/Projects/dirtools/dirwrapper/dosntexist")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	t.Log(myDir.Contents)
}
