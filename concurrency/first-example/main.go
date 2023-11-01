package main

import "fmt"

func printSomething(s string) {
	fmt.Println(s)
}

func main() {
	printSomething("This is the first thing to be printed!")
	printSomething("This is the second thing to be printed!")
}
