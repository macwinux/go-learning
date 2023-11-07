package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var seatingCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	//seed the generator
	rand.New(rand.NewSource(time.Now().UnixNano()))

	color.Yellow("The Sleeping Barber Problem")
	color.Yellow("---------------------------")

	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan:     clientChan,
		BarbersDoneChan: doneChan,
		Open:            true,
	}

	color.Green("The sho pis open for the day!")

	shop.addBarber("Paco")
	shop.addBarber("Juan")
	shop.addBarber("Raul")
	shop.addBarber("Amancio")
	shop.addBarber("Chiquito")
	shop.addBarber("Carlitos")
	shop.addBarber("Lorito")

	//start the barbershop as a goroutine
	// shopclosing channel is for don't accept more clients once the shop is closing
	shopClosing := make(chan bool)
	// stop the main function
	closed := make(chan bool)

	go func() {
		// wait 10 seconds
		<-time.After(timeOpen)
		// and now close the shop
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	}()

	// add clients
	i := 1
	go func() {
		for {
			//get a random number with average arrival rate
			randomMillseconds := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMillseconds)):
				shop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()

	// block until the barbershop is closed
	<-closed
}
