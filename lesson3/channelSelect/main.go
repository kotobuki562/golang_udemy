package main

import "fmt"

// Channel
// select

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)

	// これらがあるとランダムに出力されてしまう
	ch2 <- "A"
	ch1 <- 1
	ch2 <- "B"
	ch1 <- 2

	// どちらかのチャネルをselectして行えばデッドロックされない
	// しかし、上のv1のみ計算される。ランダムになってしまう
	select {
	case v1 := <- ch1:
		fmt.Println(v1 + 1000)
	case v2 := <- ch2:
		fmt.Println(v2 + "!?")
	default:
		fmt.Println("どちらでもない")
	}

	ch3 := make(chan int)
	ch4 := make(chan int)
	ch5 := make(chan int)

	// reciever
	go func() {
		for {
			i := <- ch3
			ch4 <- i * 2
		}
	}()

	go func() {
		for {
			i2 := <- ch4
			ch5 <- i2 - 1
		}
	}()

	n := 0
	// Lを定義しないとbreakしても無限ループ
	L:
	for {
		select {
			case ch3 <- n:
				n++
			case i3 := <- ch5:
				fmt.Println("recieve", i3)
			default:
				if n > 100 {
					break L
				}
		}
		// if n > 100 {
		// 	break
		// }
	}
	/*
	recieve -1
	recieve 0
	recieve 1
	...
	recieve 197
	*/
}