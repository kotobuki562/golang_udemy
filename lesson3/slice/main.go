package main

import (
	"fmt"
)

// スライス
// 宣言・操作

func main() {
	// 配列は[]内に要素数を指定していた
	var sl []int
	fmt.Println(sl)
	// []

	var sl2 []int = []int{100,200}
	fmt.Println(sl2)
	// [100 200]

	sl3 := []string{"A", "B"}
	fmt.Println(sl3)
	// [A B]

	sl4 := make([]int, 5)
	fmt.Println(sl4)
	// [0 0 0 0 0]

	sl2[0] = 1000
	fmt.Println(sl2)
	// [1000 200]

	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5)

	fmt.Println(sl5[0])

	// indexの2と4
	fmt.Println(sl5[2:4])
	// [3 4]

	// indexの2の手前まで
	fmt.Println(sl5[:2])
	// [1 2]

	// indexの3から最後
	fmt.Println(sl5[2:])
	// [3 4 5]

	// indexの0から最後
	fmt.Println(sl5[:])
	// [1 2 3 4 5]

	// 最後の要素
	fmt.Println(sl5[len(sl5)-1])
		
	// 一番最初と最後以外の要素
	fmt.Println(sl5[1 : len(sl5)-1])
	// [2 3 4]
}