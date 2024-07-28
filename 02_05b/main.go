package main

import (
	"fmt"
	"log"
	"sync"
)

// setup constants
const baristaCount = 3
const customerCount = 20
const maxOrderCount = 40

// coffeeShop struct holds the state of the coffee shop
type coffeeShop struct {
	orderCount int          // total number of orders processed
	orderLock  sync.Mutex   // mutex to prevent race conditions on orderCount

	orderCoffee  chan struct{} // channel for placing orders
	finishCoffee chan struct{} // channel for finishing orders
	closeShop    chan struct{} // channel to signal closing of the shop
}

// registerOrder ensures that the order made by the baristas is counted
func (p *coffeeShop) registerOrder() {
	p.orderLock.Lock()        // lock to ensure exclusive access to orderCount
	defer p.orderLock.Unlock() // defer unlock to ensure it happens after increment
	p.orderCount++
	if p.orderCount == maxOrderCount { // check if max orders are reached
		close(p.closeShop) // signal to close the shop
	}
}

// barista is the resource producer of the coffee shop
func (p *coffeeShop) barista(name string) {
	for {
		select {
		case <-p.orderCoffee: // receive an order
			p.registerOrder() // register the order
			log.Printf("%s makes a coffee.\n", name)
			p.finishCoffee <- struct{}{} // signal that coffee is finished
		case <-p.closeShop: // receive signal to close shop
			log.Printf("%s stops working. Bye!\n", name)
			return // exit the infinite loop
		}
	}
}

// customer is the resource consumer of the coffee shop
func (p *coffeeShop) customer(name string) {
	for {
		select {
		case p.orderCoffee <- struct{}{}: // place an order
			log.Printf("%s orders a coffee!", name)
			<-p.finishCoffee // wait for the coffee to be finished
			log.Printf("%s enjoys a coffee!\n", name)
		case <-p.closeShop: // receive signal to close shop
			log.Printf("%s leaves the shop! Bye!\n", name)
			return // exit the infinite loop
		}
	}
}

func main() {
	log.Println("Welcome to the Level Up Go coffee shop!")

	// channels for orders and closing the shop
	orderCoffee := make(chan struct{}, baristaCount)
	finishCoffee := make(chan struct{}, baristaCount)
	closeShop := make(chan struct{})

	// initialize the coffeeShop struct
	p := coffeeShop{
		orderCoffee:  orderCoffee,
		finishCoffee: finishCoffee,
		closeShop:    closeShop,
	}

	// create barista goroutines
	for i := 0; i < baristaCount; i++ {
		go p.barista(fmt.Sprint("Barista-", i))
	}

	// create customer goroutines
	for i := 0; i < customerCount; i++ {
		go p.customer(fmt.Sprint("Customer-", i))
	}

	// block main goroutine until closeShop channel is closed
	<-closeShop
	log.Println("The Level Up Go coffee shop has closed! Bye!")
}
