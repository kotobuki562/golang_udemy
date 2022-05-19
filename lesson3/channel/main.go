package main

import (
	"fmt"
)

// channel
// 複数のゴルーチン間でのデータの受け渡しをするために設計されたデータ構造
// 宣言、操作

func main() {
	// チャネルの宣言
	var ch1 chan int

	/*
	受信用のチャネルとして明示的に定義
	var ch2 <-chan int

	送信用のチャネルとして明示的に定義
	var ch3 chan<- int
	*/

	// makeを使ってチャネルに書き込みや読み込みを行えるようにする
	ch1 = make(chan int)

	ch2 := make(chan int)

	fmt.Println(cap(ch1))
	// 0
	fmt.Println(cap(ch2))
	// 0

	// bufサイズを第二引数で指定する
	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3))
	// 5

	// ch3に1を送信する
	ch3 <- 1
	fmt.Println(len(ch3))
	// 1

	ch3 <- 2
	ch3 <- 3
	fmt.Println("len", len(ch3))
	// 3

	i := <-ch3
	fmt.Println(i)
	// 1
	fmt.Println("len", len(ch3))
	// 2

	i2 := <-ch3
	fmt.Println(i2)
	// 2
	fmt.Println("len", len(ch3))
	// 1

	fmt.Println(<-ch3)
	// 3
	fmt.Println("len", len(ch3))
	// 0

	ch3 <- 1
	// ここで一つ取り出すことでdeadlockにはならない
	fmt.Println(<-ch3)
	// 1
	ch3 <- 2
	fmt.Println(<-ch3)
	// 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	ch3 <- 6
	// deadlock!
	// bufを超えた場合はdeadlockエラーになる

	// 先に入れたデータから順番に取り出される
}