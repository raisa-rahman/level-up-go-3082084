package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"sort"
)

const path = "items.json"

// SaleItem represents the item part of the big sale.
type SaleItem struct {
	Name           string  `json:"name"`
	OriginalPrice  float64 `json:"originalPrice"`
	ReducedPrice   float64 `json:"reducedPrice"`
	SalePercentage float64
}

// matchSales filters items within the budget, calculates the sale percentage,
// and sorts the array in descending order of sale percentage.
func matchSales(budget float64, items []SaleItem) []SaleItem {
	var mi []SaleItem
	for _, si := range items {
		if si.ReducedPrice <= budget {
			// Calculate the sale percentage
			si.SalePercentage = -(si.ReducedPrice - si.OriginalPrice) /
				si.OriginalPrice * 100
			mi = append(mi, si)
		}
	}
	// Sort the items by SalePercentage in descending order
	sort.Slice(mi, func(i, j int) bool {
		return mi[i].SalePercentage > mi[j].SalePercentage
	})

	return mi
}

func main() {
	// Parse the budget flag from the command line
	budget := flag.Float64("budget", 0.0,
		"The max budget you want to shop with.")
	flag.Parse()
	// Import items data from the JSON file
	items := importData()
	// Match items within the budget and calculate their sale percentage
	matchedItems := matchSales(*budget, items)
	// Print the matched items
	printItems(matchedItems)
}

// printItems prints the items and their sale percentages.
func printItems(items []SaleItem) {
	log.Println("The BIG sale has started with our amazing offers!")
	if len(items) == 0 {
		log.Println("No items found. :( Try increasing your budget.")
	}
	for i, r := range items {
		// Print the details of each item
		log.Printf("[%d]: %s is %.2f%% OFF! Get it now for JUST %.2f!\n",
			i, r.Name, r.SalePercentage, r.ReducedPrice)
	}
}

// importData reads the items data from the JSON file and
// unmarshals it into a slice of SaleItem.
func importData() []SaleItem {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data []SaleItem
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}