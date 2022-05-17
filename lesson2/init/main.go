package main

import "fmt"

// init
// 初期化

// initはmain()よりも先に実行される
func init() {
	fmt.Println("init")
}

// 複数init()を定義できるが基本的に一つにしておく
func init() {
	fmt.Println("init2")
}

func main() {

	fmt.Println("Main")
	// init
	// init2
  // Main
}