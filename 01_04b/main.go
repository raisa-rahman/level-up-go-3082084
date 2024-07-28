package main

import (
	"flag"
	"log"
	"math"
)

// coin contains the name and value of a coin
type coin struct {
	name  string
	value float64
}

// coins is the list of values available for making change.
// These coins are ordered from highest to lowest value.
var coins = []coin{
	{name: "1 pound", value: 1},
	{name: "50 pence", value: 0.50},
	{name: "20 pence", value: 0.20},
	{name: "10 pence", value: 0.10},
	{name: "5 pence", value: 0.05},
	{name: "1 penny", value: 0.01},
}

// calculateChange returns the coins required to make the given amount of change.
// It uses the least number of coins possible by starting with the highest value coins.
func calculateChange(amount float64) map[coin]int {
	change := make(map[coin]int) // Initialize the map to save the coins and their counts.
	for _, coin := range coins { // Loop over the coin slice, using the highest value coins first.
		if amount >= coin.value { // Check if the amount is larger or equal to the coin value.
			count := math.Floor(amount / coin.value) // Calculate the number of coins of this type.
			amount = amount - count*coin.value // Update the amount as the leftover amount.
			change[coin] = int(count) // Save the count of this coin in the change map.
		}
	}
	return change // Return the change map.
}

// printCoins prints all the coins in the map to the terminal.
func printCoins(change map[coin]int) {
	if len(change) == 0 { // If no change was calculated, print a message.
		log.Println("No change found.")
		return
	}
	log.Println("Change has been calculated.") // Print a message indicating change was calculated.
	for coin, count := range change { // Loop over the change map.
		log.Printf("%d x %s \n", count, coin.name) // Print the count and name of each coin.
	}
}

func main() {
	amount := flag.Float64("amount", 0.0, "The amount you want to make change for") // Read the amount from the command-line argument.
	flag.Parse()
	change := calculateChange(*amount) // Calculate the change for the given amount.
	printCoins(change) // Print the calculated change.
}
