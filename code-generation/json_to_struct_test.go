package code_generation_test

import (
	"fmt"
	"testing"

	"github.com/LintaoAmons/toolbox-utils/code-generation"
)

func TestJsonToStructInstance(t *testing.T) {
	a, _ := code_generation.JsonToStructInstance("TestSample", 	"Test",`{"a": "b", "c": 1, "d": true}`)
	fmt.Print(a)
	t.Log(a)
}
