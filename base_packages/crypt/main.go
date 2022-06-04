package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

//crypt

func main() {
	// MDハッシュ値を生成する
	// 任意の文字列からMDSハッシュ値を生成する処理
	h := md5.New()

	io.WriteString(h, "ABCDE")

	fmt.Println(h.Sum(nil))
	// [46 205 222 57 89 5 29 145 63 97 177 69 121 234 19 109]

	// %xで16進数の文字列を生成する
	s := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(s)
	// 8b1a9953c4611296a827abf8c47804d7
}