package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const path = "friends.json"

// Friend represents a friend and their connections.
type Friend struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Friends []string `json:"friends"`
}

// hearGossip indicates that the friend has heard the gossip.
func (f *Friend) hearGossip() {
	log.Printf("%s has heard the gossip!\n", f.Name)
}

// Friends represents the map of friends and connections
type Friends struct {
	fmap map[string]Friend
}

// getFriend fetches the friend given an id.
func (f *Friends) getFriend(id string) Friend {
	return f.fmap[id]
}

// getRandomFriend returns a random friend.
func (f *Friends) getRandomFriend() Friend {
	rand.Seed(time.Now().Unix())
	id := (rand.Intn(len(f.fmap)-1) + 1) * 100
	return f.getFriend(fmt.Sprint(id))
}

// spreadGossip ensures that all the friends in the map have heard the news
func spreadGossip(root Friend, friends Friends, visited map[string]struct{}) {
	// Loop through the immediate friends of the root friend
	for _, id := range root.Friends {
		// Check if the friend has already heard the gossip
		if _, isVisited := visited[id]; !isVisited {
			// Fetch the friend details
			f := friends.getFriend(id)
			// Mark the friend as having heard the gossip
			f.hearGossip()
			// Add the friend to the visited map
			visited[id] = struct{}{}
			// Recursively spread the gossip to this friend's friends
			spreadGossip(f, friends, visited)
		}
	}
}

func main() {
	// Import the friend data from JSON file
	friends := importData()
	// Select a random friend to start spreading the gossip
	root := friends.getRandomFriend()
	// Mark the root friend as having heard the gossip
	root.hearGossip()
	// Create a map to keep track of visited friends
	visited := make(map[string]struct{})
	visited[root.ID] = struct{}{}
	// Start spreading the gossip from the root friend
	spreadGossip(root, friends, visited)
}

// importData reads the input data from file and creates the friends map.
func importData() Friends {
	// Read the JSON file containing friend data
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the JSON data into a slice of Friend structs
	var data []Friend
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	// Create a map to store the friends with their ID as the key
	fm := make(map[string]Friend, len(data))
	for _, d := range data {
		fm[d.ID] = d
	}

	// Return the Friends struct containing the map of friends
	return Friends{
		fmap: fm,
	}
}