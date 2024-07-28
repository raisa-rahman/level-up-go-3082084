package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

const maxSeconds = 3

// Dog represents a dog with a name
type Dog struct {
	name string
}

// fetchLeash simulates the dog fetching the leash
func (d Dog) fetchLeash() {
	log.Printf("%s goes to fetch leash.\n", d.name)
	randomSleep()
	log.Printf("%s has fetched leash. Woof woof!\n", d.name)
}

// findTreats simulates the dog finding treats
func (d Dog) findTreats() {
	log.Printf("%s goes to fetch treats.\n", d.name)
	randomSleep()
	log.Printf("%s has fetched the treats. Woof woof!\n", d.name)
}

// runOutside simulates the dog running outside
func (d Dog) runOutside() {
	log.Printf("%s starts running outside.\n", d.name)
	randomSleep()
	log.Printf("%s is having fun outside. Woof woof!\n", d.name)
}

// Owner represents an owner with a name
type Owner struct {
	name string
}

// putShoesOn simulates the owner putting on shoes
func (o Owner) putShoesOn() {
	log.Printf("%s starts putting shoes on.\n", o.name)
	randomSleep()
	log.Printf("%s finishes putting shoes on.\n", o.name)
}

// findKeys simulates the owner finding keys
func (o Owner) findKeys() {
	log.Printf("%s starts looking for keys.\n", o.name)
	randomSleep()
	log.Printf("%s has found keys.\n", o.name)
}

// lockDoor simulates the owner locking the door
func (o Owner) lockDoor() {
	log.Printf("%s starts locking the door.\n", o.name)
	randomSleep()
	log.Printf("%s has locked the door.\n", o.name)
}

// randomSleep sleeps for a random duration to simulate variable task times
func randomSleep() {
	r := rand.Intn(maxSeconds)
	time.Sleep(time.Duration(r)*time.Second + 500*time.Millisecond)
}

// main function initializes the owner and dog, defines the actions, and starts the walk simulation
func main() {
	owner := Owner{name: "Jimmy"}
	dog := Dog{name: "Lucky"}
	ownerActions := []func(){
		owner.putShoesOn,
		owner.findKeys,
		owner.lockDoor,
	}
	dogActions := []func(){
		dog.fetchLeash,
		dog.findTreats,
		dog.runOutside,
	}
	executeWalk(ownerActions, dogActions)
}

// executeWalk executes the given actions for the owner and the dog concurrently
func executeWalk(ownerActions []func(), dogActions []func()) {
	var wg sync.WaitGroup
	wg.Add(2)  // Add two to the WaitGroup counter for owner and dog
	defer wg.Wait()  // Defer the Wait call to block until all actions are completed

	// Helper function to execute a list of actions
	executeActions := func(actions []func()) {
		defer wg.Done()  // Decrement the WaitGroup counter when the goroutine completes
		for _, a := range actions {
			a()  // Execute each action in order
		}
	}

	// Start owner actions in a separate goroutine
	go executeActions(ownerActions)
	// Start dog actions in a separate goroutine
	go executeActions(dogActions)
}
