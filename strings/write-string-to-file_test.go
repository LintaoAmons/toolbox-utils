package strings_test

import (
	"io/ioutil"
	"os"
	"testing"

	toolboxutils "github.com/LintaoAmons/toolbox-utils/strings"
)

func TestWriteStringToFile(t *testing.T) {
	// Create a temporary file to write to
	file, err := ioutil.TempFile("", "testfile-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	// Write a test string to the file
	data := "test string\n"
	err = toolboxutils.WriteStringToFile(file.Name(), data)
	if err != nil {
		t.Fatal(err)
	}

	// Read the contents of the file
	contents, err := ioutil.ReadFile(file.Name())
	if err != nil {
		t.Fatal(err)
	}

	// Ensure that the contents of the file match the expected string
	expected := data
	actual := string(contents)
	if expected != actual {
		t.Errorf("expected %q, but got %q", expected, actual)
	}
}
