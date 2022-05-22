package main

import (
	"fmt"
)

// Channel
// for

func main() {
	ch1 := make(chan int, 3)

	ch1 <- 1
	ch1 <- 2
	ch1 <- 3

	// closeで締めないとdeadrockになる
	close(ch1)
	for i := range ch1 {
		fmt.Println(i)
	}
	/*
	1
	2
	3
	fatal error: all goroutines are asleep - deadlock!
	*/

}