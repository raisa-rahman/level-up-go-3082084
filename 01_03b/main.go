package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

const path = "entries.json"

// raffleEntry is the struct we unmarshal raffle entries into
// It represents each entry in the raffle with an ID and Name.
type raffleEntry struct {
	ID   string `json:"id"`   // ID field mapped to the JSON "id"
	Name string `json:"name"` // Name field mapped to the JSON "name"
}

// importData reads the raffle entries from file and creates the entries slice.
// It reads the JSON data from a file, unmarshals it into a slice of raffleEntry structs,
// and returns this slice. If an error occurs during reading or unmarshaling, the program
// will log the error and terminate.
func importData() []raffleEntry {
	// Read the file specified by the path constant
	file, err := os.ReadFile(path)
	if err != nil {
		// Log the error and terminate the program if the file cannot be read
		log.Fatal(err)
	}

	var data []raffleEntry
	// Unmarshal the JSON data into the data slice
	err = json.Unmarshal(file, &data)
	if err != nil {
		// Log the error and terminate the program if unmarshaling fails
		log.Fatal(err)
	}

	// Return the slice of raffle entries
	return data
}

// getWinner returns a random winner from a slice of raffle entries.
// It uses the rand package to generate a pseudo-random number based on the current time,
// selects a winner from the entries slice, and returns this winner.
func getWinner(entries []raffleEntry) raffleEntry {
	// Seed the random number generator with the current Unix time
	rand.Seed(time.Now().Unix())
	// Generate a random index based on the length of the entries slice
	wi := rand.Intn(len(entries))
	// Return the entry at the randomly generated index
	return entries[wi]
}

func main() {
	// Import the raffle entries from the JSON file
	entries := importData()
	log.Println("And... the raffle winning entry is...")
	// Get a random winner from the imported entries
	winner := getWinner(entries)
	// Pause for dramatic effect
	time.Sleep(500 * time.Millisecond)
	// Log the winning entry
	log.Println(winner)
}