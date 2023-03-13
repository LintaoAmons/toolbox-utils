package toolboxutils_test

import (
	"fmt"
	"testing"
	"time"

	toolboxutils "github.com/LintaoAmons/toolbox-utils"
)

func TestStructToStringWithTime(t *testing.T) {
	// Define a test struct that includes a time.Time field
	type MyStruct struct {
		ID        int
		Name      string
		CreatedAt time.Time
	}

	// Create a test struct instance with a time.Time value
	createdAt := time.Date(2022, 3, 13, 12, 0, 0, 0, time.UTC)
	s := MyStruct{
		ID:        123,
		Name:      "Test",
		CreatedAt: createdAt,
	}

	// Call the StructToString function with the test struct
	result := toolboxutils.StructToString(s)

	// Check that the result includes the formatted time.Time field value in the correct format
	expected := fmt.Sprintf("'%d', '%s', '%s'::timestamp WITH TIME ZONE", s.ID, s.Name, createdAt.Format("2006-01-02 15:04:05.999999 -07:00"))
	if result != expected {
		t.Errorf("StructToString returned %q, expected %q", result, expected)
	}
}
