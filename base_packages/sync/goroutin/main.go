package main

import (
	"fmt"
	"sync"
)

//sync
// ゴルーチンの終了を待ち受ける
func main() {
	// sync.WateGroupを生成
	wg := new(sync.WaitGroup)
	// 待ち受けするゴルーチおんは3つ
	wg.Add(3)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("1st Goroutine")
		}
		wg.Done() //完了
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("2nd Goroutine")
		}
		wg.Done() //完了
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("3rd Goroutine")
		}
		wg.Done() //完了
	}()

	// ゴルーチンの完了を待ち受ける
	// Doneが3つ完了するまで待つ
	wg.Wait()


	// WateGroupがなければ何も出力されない
}