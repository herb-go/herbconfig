package source

import "testing"
import "html"

func TestText(t *testing.T) {
	v := "http://<>"
	msg := "text://" + html.UnescapeString(v)
	text, err := New(msg)
	if err != nil {
		t.Fatal(err)
	}
	bs, err := Read(text)
	if err != nil {
		t.Fatal(err)
	}
	if string(bs) != v {
		t.Fatal(string(bs))
	}
	if text.AbsolutePath() != "" {
		t.Fatal(text)
	}
}
