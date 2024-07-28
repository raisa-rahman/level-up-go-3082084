package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// the amount of bidders we have at our auction
const bidderCount = 10

// initial wallet value for all bidders
const walletAmount = 250

// items is the map of auction items
var items = []string{
	"The \"Best Gopher\" trophy",
	"The \"Learn Go with Adelina\" experience",
	"Two tickets to a Go conference",
	"Signed copy of \"Beautiful Go code\"",
	"Vintage Gopher plushie",
}

// bid is a type that pairs the bidder id and the amount they want to bid
type bid struct {
	bidderID string
	amount   int
}

// auctioneer receives bids and announces winners
type auctioneer struct {
	bidders map[string]*bidder
}

// runAuction manages the auction for all the items to be sold
func (a *auctioneer) runAuction(bids <-chan bid, open chan<- struct{}) {
	for _, item := range items {
		log.Printf("Opening bids for %s!\n", item)
		a.openBids(open)             // Signal to all bidders to open bids
		a.processWinner(item, bids)  // Process and announce the winner for the current item
	}
}

// openBids signals to all bidders that the auction for the item is open.
func (a *auctioneer) openBids(open chan<- struct{}) {
	for i := 0; i < bidderCount; i++ {
		open <- struct{}{}  // Send an open signal to each bidder
	}
}

// processWinner reads all the bids and finds the winner of the auction
func (a *auctioneer) processWinner(item string, bids <-chan bid) {
	var winner bid
	for i := 0; i < bidderCount; i++ {
		cb := <-bids  // Collect each bid from the channel
		if winner.amount < cb.amount {
			winner = cb  // Update the winner if the current bid is higher
		}
	}
	log.Printf("%s is sold to %s for %d!\n", item,
		winner.bidderID, winner.amount)
	a.bidders[winner.bidderID].payBid(winner.amount)  // Instruct the winning bidder to pay
}

// bidder is a type that represents a bidder with their wallet amount
type bidder struct {
	id     string
	wallet int
}

// placeBid generates a random amount and places it on the bids channel
func (b *bidder) placeBid(out chan<- bid, open <-chan struct{}) {
	for i := 0; i < len(items); i++ {
		<-open  // Wait for the auctioneer to open bids for the current item
		currentBid := bid{bidderID: b.id}
		if b.wallet > 0 {
			currentBid.amount = getRandomAmount(b.wallet)  // Generate a random bid amount
		}
		out <- currentBid  // Send the bid to the auctioneer
	}
}

// payBid subtracts the bid amount from the wallet of the auction winner
func (b *bidder) payBid(amount int) {
	b.wallet -= amount  // Deduct the winning bid amount from the bidder's wallet
}

func main() {
	rand.Seed(time.Now().UnixNano())
	log.Println("Welcome to the LinkedIn Learning auction.")
	bidders := make(map[string]*bidder, bidderCount)
	bids := make(chan bid, bidderCount)
	openBids := make(chan struct{})
	for i := 0; i < bidderCount; i++ {
		id := fmt.Sprint("Bidder ", i)
		b := bidder{
			id:     id,
			wallet: walletAmount,
		}
		bidders[id] = &b
		go b.placeBid(bids, openBids)  // Start a goroutine for each bidder to place bids
	}
	a := auctioneer{
		bidders: bidders,
	}
	a.runAuction(bids, openBids)  // Run the auction
	log.Println("The LinkedIn Learning auction has finished!")
}

// getRandomAmount generates a random integer amount up to max
func getRandomAmount(max int) int {
	return rand.Intn(int(max))
}