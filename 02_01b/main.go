package main

import (
	"flag"
	"log"
)

// Define a slice of messages to be printed
var messages = []string{
	"Hello!",
	"How are you?",
	"Are you just going to repeat what I say?",
	"So immature",
	"Stop copying me!",
}

// repeat concurrently prints out the given message n times
func repeat(n int, message string) {
	// Create a channel to signal completion of goroutines
	ch := make(chan struct{})
	
	// Start n goroutines to print the message concurrently
	for i := 0; i < n; i++ {
		go func(i int) {
			// Log the message along with the goroutine index
			log.Printf("[G%d]:%s\n", i, message)
			// Send a signal to the channel indicating this goroutine is done
			ch <- struct{}{}
		}(i)
	}

	// Block the main goroutine until all n goroutines have signaled completion
	for i := 0; i < n; i++ {
		<-ch
	}
}

func main() {
	// Parse the fan-out factor from command line arguments
	factor := flag.Int64("factor", 0, "The fan-out factor to repeat by")
	flag.Parse()

	// Loop through each message in the messages slice
	for _, m := range messages {
		// Log the message from the main goroutine
		log.Printf("[Main]:%s\n", m)
		// Call the repeat function to print the message n times concurrently
		repeat(int(*factor), m)
	}
}