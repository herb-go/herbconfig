package configuration

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	file, err := ioutil.TempFile("", "herb-go-test")
	if err != nil {
		t.Fatal(err)
	}
	_, err = file.WriteString("testcontent")
	if err != nil {
		t.Fatal(err)
	}
	name := file.Name()
	file.Close()
	defer os.Remove(name)
	file1 := File(name)
	data, err := Read(file1)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "testcontent" {
		t.Fatal(string(data))
	}

	if file1.AbsolutePath() == "" {
		t.Fatal(file1.AbsolutePath())
	}
	file2 := File(name + ".notexists")

	if IsSame(file1, file2) {
		t.Fatal(file2.ID())
	}
}
