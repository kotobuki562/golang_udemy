package main

import (
	"fmt"
	"time"
)

// go goroutin(ゴルーチン)
// 並行処理

func sub() {
	for {
		fmt.Println("sub loop")
		// 100ms感覚で実行される
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// 「go」を付けることで並行処理できる
	// 「go」がなければsub()が無限ループして永遠に下のforに到達しない
	go sub()
	go sub()

	for {
		fmt.Println("main loop")
		time.Sleep(200 * time.Millisecond)
	}
}