package main

import (
	"fmt"
	"golang_udemy/lesson5/test/alib"
)

// テスト

// パスを指定してテストを実行する
// go test ./alib
// go test -v ./alib

// ディレクトリ内の全てのテストを実行する
// go test -v ./...

// パッケージのcover率が出力される
// go test -cover ./...

func IsOne(i int) bool {
	if i == 1 {
		return true
	}else{
		return false
	}
}

func main() {
	fmt.Println(IsOne(1))
	fmt.Println(IsOne(2))

	s := []int{1, 2, 3, 4, 5}
	fmt.Println(alib.Average(s))
	// 3
}