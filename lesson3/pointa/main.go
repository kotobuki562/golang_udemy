package main

import (
	"fmt"
)

// ポインタ

func Double(i int) {
	i = i * 2
}

// 参照型
func DoubleV2(i *int) {
	*i = *i * 2
}

// 参照型
// 参照型であれば値の更新にポインタは必要ない
func DoubleV3(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}

func main() {
	var n int = 100
	fmt.Println(n)

	// &を付けるとメモリ上のアドレスが出力される
	fmt.Println(&n)
	// 0xc0000b2008

	Double(n)
	fmt.Println(n)
	// 100

	// アドレスを渡している
	// nの更新ができるようになる
	var p *int = &n
	/*
	fmt.Println(p)
	// 0xc0000b2008
	fmt.Println(*p)
	// 100

	*p = 300
	fmt.Println(n)
	// 200

	n = 200
	fmt.Println(*p)
	// 200
	*/
	
	DoubleV2(&n)
	fmt.Println(n)
	// 200

	DoubleV2(p)
	fmt.Println(*p)
	// 400

	var sl []int = []int{1, 2, 3}
	DoubleV3(sl)
	fmt.Println(sl)
	// [2 4 6]
}