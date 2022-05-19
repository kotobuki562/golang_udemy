package main

import (
	"fmt"
)

// スライス
// copy

func main() {
	sl := []int{100, 200}
	sl2 := sl

	// 元データのslに1000が入る
	sl2[0] = 1000
	fmt.Println(sl)
	// [1000 200]

	// 基本型の場合はそれぞれ独立している
	// 山椒型の場合は元データも更新される
	var i int = 10
	i2 := i
	i2 = 100
	fmt.Println(i)
	// 10
	fmt.Println(i2)
	// 100

	sl3 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl3)
	sl4 := make([]int, 5, 10)
	fmt.Println(sl4)
	// 第一引数にコピー先、第二引数にコピー元
	// 先頭から塗りつぶすようにコピーしていくしよう
	n := copy(sl4 ,sl3)
	fmt.Println(n, sl4)
	// 5 [1 2 3 4 5]
}