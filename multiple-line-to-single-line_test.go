package toolboxutils_test

import (
	"testing"

	toolboxutils "github.com/LintaoAmons/toolbox-utils"
)

func TestMultilineToSingleline(t *testing.T) {
	// Test input and expected output
	input := "This is a multi-line\nstring with\nthree lines."
	expected := "This is a multi-line string with three lines."

	// Call the function being tested
	output := toolboxutils.MultilineToSingleline(input)

	// Check if the output matches the expected value
	if output != expected {
		t.Errorf("expected %q, but got %q", expected, output)
	}
}
