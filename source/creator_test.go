package source

import "testing"

func TestFileCreator(t *testing.T) {
	tp := getFileTypeFromID("http://www.example.com")
	if tp != "http" {
		t.Fatal(tp)
	}
	tp = getFileTypeFromID("www.example.com")
	if tp != "" {
		t.Fatal(tp)
	}
	_, err := New("notexistscheme://123")
	if err == nil || GetErrorType(err) != ErrTypeFileObjectSchemeNotavaliable {
		t.Fatal(err)
	}
	_, err = New("")
	if err == nil || GetErrorType(err) != ErrTypeFileObjectSchemeNotavaliable {
		t.Fatal(err)
	}
	localfile := File("/tmp/test.go")
	file, err := New(localfile.ID())
	if !IsSame(localfile, file) {
		t.Fatal(file.ID())
	}

	textfile := Text("test")
	file, err = New(textfile.ID())
	if !IsSame(textfile, file) {
		t.Fatal(file.ID())
	}
}
