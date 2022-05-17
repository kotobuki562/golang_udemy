package main

import (
	"fmt"
	"strconv"
)

// if
// 条件分岐
// エラーハンドリング

func main() {
	var s string = "ABC"

	i, error := strconv.Atoi(s)
	if error != nil {
		fmt.Println(error)
		// strconv.Atoi: parsing "ABC": invalid syntax
	}
	fmt.Printf("i = %T\n", i)
	// i = int
}