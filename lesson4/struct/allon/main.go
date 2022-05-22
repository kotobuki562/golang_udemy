package main

import "fmt"

// struct
// 独自型

type MyInt int

func (mi MyInt) Print() {
	fmt.Println(mi)
}

func main() {
	var mi MyInt
	mi = 10
	fmt.Println(mi)
	// 10
	fmt.Printf("%T\n", mi)
	// main.MyInt

	// i := 100
	// fmt.Println(mi + i)
	// error

	mi.Print()
	// 10

}