package main

//
import (
	// 通常のパッケージ宣言
	"fmt"
	// 名前空間を指定したパッケージ宣言
	// あまり水晶はされていない
	. "fmt"
	// 名前を指定してパッケージ宣言
	f "fmt"

	// gopathはいかになければエラーになる
	"golang_udemy/lesson5/public/foo"
)

// スコープ

func appName() string {
	const appName = "GoApp"
	var Version string = "1.0"
	return appName + " " + Version
}

// string型のbを返す
func Do(s string) (b string) {
	// これだとスコープhが重複する
	// var b string = s
	b = s
	{
		b := "BBBB"
		f.Println(b)
		// BBBB
	}
	f.Println(b)
	// Hello
	return b
}

func main() {
	f.Println(foo.Max)
	// 1

	// fmt.Println(foo.min)
	// cannot refer to unexported name foo.min

	// 関数を呼び出しているのでminも参照可能
	fmt.Println(foo.ReturnMin())
	// 1

	Println(foo.Max)

	fmt.Println(appName())
	// GoApp 1.0

	// f.Println(AppName, Version)
	// ./main.go:41:12: undefined: AppName
  // ./main.go:41:21: undefined: Version

	f.Println(Do("Hello"))
	// error
}