package main

import (
	"encoding/json"
	"log"
	"os"
)

// User represents a user record.
type User struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

const path = "users.json"

// getBiggestMarket takes in the slice of users and
// returns the biggest market (the country with the most users).
func getBiggestMarket(users []User) (string, int) {
	// Initialize a map to count the number of users per country.
	counts := make(map[string]int)
	// Loop over each user and increment the count for their country.
	for _, u := range users {
		counts[u.Country]++
	}

	// Variables to store the country with the maximum users and the user count.
	maxCountry := ""
	maxCount := 0
	// Iterate over the map to find the country with the highest user count.
	for country, count := range counts {
		if count > maxCount {
			maxCount = count
			maxCountry = country
		}
	}

	// Return the country with the most users and the user count.
	return maxCountry, maxCount
}

func main() {
	// Import user data from the JSON file.
	users := importData()
	// Get the country with the most users and the user count.
	country, count := getBiggestMarket(users)
	// Log the result.
	log.Printf("The biggest user market is %s with %d users.\n", country, count)
}

// importData reads the user entries from the JSON file and
// returns a slice of User structs.
func importData() []User {
	// Read the JSON file.
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the JSON data into a slice of User structs.
	var data []User
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	// Return the slice of User structs.
	return data
}
