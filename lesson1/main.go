package main

import (
	"fmt"
)

// interface

func main() {
	var x interface{}
	fmt.Println(x)
	// <nil>
	// 値を何も持っていない状態であることを表す

	x = 1
	fmt.Println(x)
	// 1
	x = "hello"
	fmt.Println(x)

	x = [3]int{1,2,3}
	fmt.Println(x)

	x = 3.14
	fmt.Println(x)

	// x = 2
	// fmt.Println(x + 1)
	//  invalid operation: x + 1 (mismatched types interface {} and int)
	// interfaceは型の互換性はあるが他の方と併用して使えない
}