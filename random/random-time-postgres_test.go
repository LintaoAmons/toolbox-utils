package random_test

import (
	"testing"

	"github.com/LintaoAmons/toolbox-utils/random"
)

func TestRandomTimePostgres(t *testing.T) {
	startDate := "2022-12-01"
	endDate := "2022-12-31"

	// Generate 10 random dates within the range
	for i := 0; i < 10; i++ {
		dateStr, err := random.RandomTimePostgres(startDate, endDate)
		if err != nil {
			t.Errorf("Error generating random date: %s", err)
		}
		t.Logf("Generated date: %s", dateStr)
	}
}
