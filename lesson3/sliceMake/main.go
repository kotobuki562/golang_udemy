package main

import (
	"fmt"
)

// スライス
// append make len cap

func main() {
	sl := []int{100, 200}
	fmt.Println(sl)

	// 300という要素を追加する
	sl = append(sl, 300)
	fmt.Println(sl)
	// [100 200 300]
	
	sl = append(sl, 400, 500)
	fmt.Println(sl)
	// [100 200 300 400 500]

	sl2 := make([]int, 5)
	fmt.Println(sl2)

	fmt.Println(len(sl2))
	// 5

	// sl2の容量を測る
	fmt.Println(cap(sl2))
	// 5

	// 第二引数で容量を指定する
	sl3 := make([]int, 5, 10)

	fmt.Println(sl3)
	// [0 0 0 0 0]

	fmt.Println(len(sl3))
	// 5

	fmt.Println(cap(sl3))
	// 10

	sl3 = append(sl3, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(sl3)
	// [0 0 0 0 0 1 2 3 4 5 6 7]

	fmt.Println(cap(sl3))
	// 20

}