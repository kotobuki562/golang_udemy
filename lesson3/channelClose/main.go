package main

import (
	"fmt"
	"time"
)

// Channel
// close

func recieve(name string, ch chan int) {
	for {
		// 受け取ったchがclose状態かどうかをチェック
		i, ok := <- ch
		// もしclose状態ならループを抜ける
		if !ok {
			break
		}
		// closeでなければ受け取った値を表示
		fmt.Println(name, i)
	}
	fmt.Println(name, "end")
}

func main() {
	// 最初はopen状態
	ch1 := make(chan int, 2)


	// ch1 <- 1

	// close(ch1)

	// ch1 <- 1
	// runtimeError

	// fmt.Println(ch1)
	// 0xc0000b6000

	// runtimeエラーにならない
	// ch1がopen状態であればtrueになる
	// i, ok := <- ch1
	// fmt.Println(i, ok)
	// 0 false

	// i2, ok := <- ch1
	// fmt.Println(i2, ok)
	// 0 false

	go recieve("1.goroutin", ch1)
	go recieve("2.goroutin", ch1)
	go recieve("3.goroutin", ch1)

	i := 0
	for i < 100 {
		ch1 <- i
		i++
	}

	close(ch1)
	time.Sleep(3 + time.Second)


}