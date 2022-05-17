package main

import (
	"fmt"
	"os"
)

// defer
// 関数の終了時に登録することができる

func TestDefer() {
	// deferは関数の終了時に登録することができる
	defer fmt.Println("END")
	fmt.Println("START")
}

func Rundefer() {
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	defer fmt.Println("defer3")
}

func main() {
	TestDefer()
	// START
  //	END

	// 複数のderferを登録することもできる
	defer func() {
		fmt.Println("defer1")
		fmt.Println("defer2")
		fmt.Println("defer2")
	}()

	// 後に実行した関数から実行される
	Rundefer()
	// defer3
  // defer2
  // defer1

	file, error := os.Create("test.txt")
	if error != nil {
		fmt.Println(error)
	}
	// test.txtというファイルがcurdirに作成される
	// deferでファイルを閉じる
	defer file.Close()

	file.Write([]byte("test"))

}