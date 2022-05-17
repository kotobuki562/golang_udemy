package main

import "fmt"

// ジェネレーター
// クロージャーを応用する

func integers() func() int {
	i := 0
	// iに対して増分して返す関数
	return func() int {
		i++
		return i
	}
}

func main() {
	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	// 1
	// 2
	// 3
	// 4

	// 再定義してもintsと同じ挙動をする
	others := integers()
	fmt.Println(others())
	fmt.Println(others())
	fmt.Println(others())
	fmt.Println(others())
	// 1
	// 2
	// 3
	// 4
}