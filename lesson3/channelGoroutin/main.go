package main

import (
	"fmt"
	"time"
)

// channel
// チャネルとゴルーチン

func reciever(c chan int) {
	for {
		i := <- c
		fmt.Println(i)
	}
} 

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	
	// deadlock
	// ゴルーチン間での送受信が前提なのでチャネル単体ではdeadlockになる

	// チャネルに値が入るのを待つ
	go reciever(ch1)
	go reciever(ch2)

	// i(int)をforでチャネルに送信
	i := 0
	for i < 100 {
		ch1 <- i
		ch2 <- i
		time.Sleep(50 * time.Millisecond)
		i++
	}
	// 1
	// 1
	//1
	// 2
	// ...
	// 99
	// 99
}