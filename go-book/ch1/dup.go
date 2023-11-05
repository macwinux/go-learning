package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// esto es igual que:
		// line := input.Text()
		// counts[line] = counts[line] +1
		text := input.Text()
		if text == "quit" {
			for line, n := range counts {
				if n > 1 {
					fmt.Printf("%d\t%s\n", n, line)
				}
			}
			return
		}
		counts[text]++
	}
}
