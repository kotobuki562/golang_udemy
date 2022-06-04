package main

import (
	"fmt"
	"sync"
	"time"
)

//sync
// 同期処理
// ミューテックスによる同期処理

/*
非同期処理における「レースコンディション」
並列動作する複数のプロセスやリソースが、同一のリソースに同時にアクセスしたら、
予想外の挙動を起こす
*/

var st struct{A, B, C int}

func UpdateAndPrint(n int) {
	st.A = n
	time.Sleep(time.Microsecond)
	st.B = n
	time.Sleep(time.Microsecond)
	st.C = n
	time.Sleep(time.Microsecond)
	fmt.Println(st)
	// {999 999 999}
	// {996 999 999}
	// {996 996 996}
	// {997 997 997}
	// このように同一の数値でないものが肺いているケースが起こってします
}

// ミューテックスを保持するパッケージ変数
var mutex *sync.Mutex

func UpdateAndPrintMutex(n int) {
	// ロックされている間は一つのgoroutinしか実行されない
	mutex.Lock()

	st.A = n
	time.Sleep(time.Microsecond)
	st.B = n
	time.Sleep(time.Microsecond)
	st.C = n
	time.Sleep(time.Microsecond)
	fmt.Println(st)
	// {995 995 995}
	// {996 996 996}
	// {997 997 997}
	// {998 998 998}
	// ちゃんと同一のデータが入る

	// 処理を終えたらロックを解放する
	mutex.Unlock()
}

func main() {
	mutex = new(sync.Mutex)
	for i := 0; i < 5; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				UpdateAndPrintMutex(i)
			}
		}()
	}
	for {

	}


}