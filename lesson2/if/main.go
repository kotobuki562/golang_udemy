package main

import "fmt"

// if
// 条件分岐

func main() {
	a := 0
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("I don,t know")
	}

	// 条件式に簡易文を使う
	if b := 100; b == 100 {
		fmt.Println("100")
		// 100
	} else {
		fmt.Println("not 100")
	}

	// if文内では2になる
	x := 0
	if x := 2; true {
		fmt.Println(x)
		// 2
	}
	fmt.Println(x)
	// 0
}