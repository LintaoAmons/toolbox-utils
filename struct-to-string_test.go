package toolboxutils_test

import (
	"testing"

	toolboxutils "github.com/LintaoAmons/toolbox-utils"
)

type TestStruct struct {
	Field1 string
	Field2 string
	Field3 string
}

func TestStructToString(t *testing.T) {
	TestStruct := newTestStruct()

	result := toolboxutils.StructToString(TestStruct)

	expected := "'field1', 'field2', 'field3'"

	if result != expected {
		t.Errorf("unexpected result: %s, expected: %s", result, expected)
	}
}

func newTestStruct() TestStruct {
	return TestStruct{
		Field1: "field1",
		Field2: "field2",
		Field3: "field3",
	}
}
