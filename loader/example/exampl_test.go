package example

import "testing"

func TestExample(t *testing.T) {
	if !ExampleData.Equal(ExampleData) {
		t.Fatal(ExampleData)
	}
	if ExampleData.Equal(&ExampleStruct{}) {
		t.Fatal(ExampleData)
	}
}
