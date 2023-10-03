package main

import (
	"fmt"
	"math/rand"
	"time"
)

func testChannels() {
	/*
		// declaring a channel that can hold a single integer value
		var myChannel = make(chan int)
		myChannel <- 1 // this adds the value 1 to the channel
		// get the value from the channel and assign it to a vairable
		var n = <- myChannel // reads from the channel
		fmt.Println(n)
	*/
	//, sending to a channel (myChannel <- 1) and receiving from the same channel (var n = <-myChannel) within the same goroutine will cause a deadlock because the goroutine is trying to both send and receive from the channel without another goroutine involved to perform the complementary operation.

	var myChannel = make(chan int)
	go processing(myChannel)
	for i := range myChannel {
		fmt.Println(i)
	}

	fmt.Println("buffers channels ...")
	var buffersChannel = make(chan int, 5)
	go processing(buffersChannel)
	for i := range buffersChannel {
		fmt.Println(i)
	}
}

func processing(c chan int) {
	defer close(c) // close the channel to inform all routines using this chanel that we are done and break the the reads from that channel
	for i := 0; i <= 5; i++ {
		c <- i
	}
	fmt.Println("processing func exits")
}

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func main() {
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"walmart", "costco.com", "wholefoods.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkToFuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)

}

// check
func checkToFuPrices(website string, chickenChannel chan string) {

	for {
		time.Sleep(time.Second)
		var tofuPrice = rand.Float32() * 20
		if tofuPrice <= MAX_TOFU_PRICE {
			chickenChannel <- website
			break
		}

	}

}

// chiken price
func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second) // pause the execution for 1 second
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
	case website := <-chickenChannel:
		fmt.Printf("\n Text sent: deal on chiken at %v", website)
	case website := <-tofuChannel:
		fmt.Println("\n Found a deal on tofu at %v", website)
	}
}
