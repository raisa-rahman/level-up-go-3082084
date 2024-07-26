package main

import (
	"flag"
	"log"
	"time"
)

var expectedFormat = "2006-01-02"

// parseTime validates and parses a given date string.
func parseTime(target string) time.Time {
	// We begin by parsing the time given the expected format.
	pt, err := time.Parse(expectedFormat, target)
	// Add an 'if' check for a non-nil parsing error or if time.Now is after the parsed time.
	if err != nil || time.Now().After(pt) {
		// In the error case, we will terminate the program.
		log.Fatal("invalid target date: ", target)
	}

	return pt
}

// calcSleeps returns the number of sleeps until the target.
func calcSleeps(target time.Time) float64 {
	// Use time.Until to get the duration until the target time.
	// Convert the duration to hours and then divide by 24 to get the number of days.
	return time.Until(target).Hours() / 24
}

func main() {
	// Read the target birthday as a command line argument.
	bday := flag.String("bday", "", "Your next bday in YYYY-MM-DD format")
	flag.Parse()
	// Parse the target birthday.
	target := parseTime(*bday)
	// Log the number of sleeps until the target birthday.
	log.Printf("You have %d sleeps until your birthday. Hurray!",
		int(calcSleeps(target)))
}