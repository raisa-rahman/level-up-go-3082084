package main

import (
	"log"
	"strings"
	"time"
)

const delay = 700 * time.Millisecond

// print outputs a message and then sleeps for a pre-determined amount
// (This print method will be used in the slowDown function to print 
// each prolonged word and then pause before the next word.)
func print(msg string) {
	log.Println(msg)
	time.Sleep(delay)
}

// slowDown takes the given string and repeats its characters
// according to their index in the string.
// (First, we use strings.Split to split the message by spaces to get a slice of words.)
func slowDown(msg string) {
	words := strings.Split(msg, " ")
	// (Range over the words to process each word individually.)
	for _, w := range words {
		var pw []string
		// (For each letter in the word, repeat it based on its index + 1.)
		for i, c := range w {
			// (strings.Repeat repeats the character c, i+1 times.)
			rb := strings.Repeat(string(c), i+1)
			// (Append the repeated block to the prolonged word slice.)
			pw = append(pw, rb)
		}
		// (Join the prolonged word slice into a single string and print it.)
		print(strings.Join(pw, ""))
	}
}

func main() {
	msg := "Time to learn about Go strings!"
	// (Invoke the slowDown function with a predetermined message.)
	slowDown(msg)
}