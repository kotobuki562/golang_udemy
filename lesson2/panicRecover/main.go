package main

import "fmt"

// panic & recover
// 例外処理(runtimeを強制的に停止させる)
// あまり使わない

func main() {
	defer func() {
		// panicが発生した時に実行されるようにするためdeferでrecoverを定義する方がいい
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	panic("runtime errpr")
	fmt.Println("one or two")
}