package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(s)
}

func main() {

	// don't want to wait sleeping, we use WaitGroup instead
	var wg sync.WaitGroup

	// go routines has no order

	words := []string{
		"alpha",
		"beta",
		"gamma",
		"delta",
		"pi",
		"eta",
		"zeta",
	}

	wg.Add(len(words))

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d : %s", i, x), &wg)
	}

	wg.Wait()
	wg.Add(1)
	printSomething("This is the second thing to be printed!", &wg)

}

