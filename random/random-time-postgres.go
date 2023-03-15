package random

import (
	"math/rand"
	"time"
)

func RandomTimePostgres(startDate, endDate string) (string, error) {

	// Parse the start and end dates as time.Time values
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return "", err
	}
	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return "", err
	}

	// Generate a random time between the start and end dates
	delta := end.Sub(start)
	rand.Seed(time.Now().UnixNano())
	randomDelta := time.Duration(rand.Int63n(int64(delta)))
	randomTime := start.Add(randomDelta)

	// Format the random time as a string
	return randomTime.Format("2006-01-02 15:04:05.000000 -07:00"), nil
}
