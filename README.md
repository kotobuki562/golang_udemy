## 【go run】

Go プログラムの手軽な実行方法

Go はコンパイルを基本とした言語ですが、直接プログラムを実行できる go run コマンドを備えています。

`go run main.go`

## 【go build】

go build コマンドは、OP で与えた Go ファイルを実行ファイル形式にコンパイルする。

`go build -o main main.go`

-o 　 OP を使用して出力する実行ファイルのファイル名を指定できる。

上記では、main.go を main としてコンパイル。

## 【コンパイルするメリット】

1. ネイティブコードにコンパイルされる

Go はネイティブコードにコンパイルされた上で実行されるので、一般的なスクリプト言語の実行速度より１０〜１００倍という高いパフォーマンスを発揮する。

2. マルチプラットフォームで動作する

Go は OS や CPU による実行環境の差を隠蔽してくれる。

この為、実行されるプラットフォームの差に気を配らなくても良い。

各実行環境で動作するプログラムを、１つのコンパイル環境から生成できる、クロスコンパイル機能を備えている。

特にこのクロスコンパイルの恩恵が大きいです。

Go は、 1 つのソースコードから様々な OS 向けのバイナリを生成するクロスコンパイルをサポートしています。

Mac で .exe ファイルを生成して Windows ユーザにそれを渡せば、受け取った瞬間にすぐに実行できるという点が非常に便利なため、

コマンドラインツールのようなものをクロスコンパイルして配布するなどといった用途で、 Go が導入される場面も目立ってきました。

環境を指定したい場合は、 GOOS と GOARCH という環境変数を先ほどのリストにある組み合わせで指定します。

Linux 環境を指定してコンパイル

`GOOS=linux GOARCH=amd64 go build -o main main.go`

## 【マルチプラットフォーム】

前述の通り、コンパイルすることで、異なる環境でも同一に動作させることが可能です。

例えば、Python などだと、実装環境と、実行環境で差がある場合動作しないことがあります。

(ホスト A で言語がインストールされているが、B ではインストールされていない場合は動作しない)

Go の場合は、build して出来上がったバイナリをそのまま配布するだけで動作するというメリットがあります。

あとは、ターミナル上で

`./main`

と実行するだけで動作します

# 変数定義

```main.go
package main

import (
	"fmt"
)

var i5 int = 500
// 明示的に変数定義ができれば関数街で定義できる。なるべく明示的に変数定義をする
// i5 := 500

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {
	/*
	明示的な変数定義
	var name string = "golang"
	*/
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t,f bool = true ,false
	fmt.Println(t,f)

	var(
		i2 int = 200
		s2 string = "Hello Go2"
	)
	fmt.Println(i2,s2)

	var i3 int
	var s3 string
	fmt.Println(i3,s3)

	i3 = 300
	s3 = "Hello Go3"
	fmt.Println(i3,s3)

	i = 150
	fmt.Println(i)

	/*
	暗黙的な変数定義
	name := "golang"
	*/

	i4 := 400
	fmt.Println(i4)
	i4 = 450
	fmt.Println(i4)
	// :=で同じ変数名を再定義できない
	// i4 := 500
	// fmt.Println(i4)

	// int型ので宣言した変数はint以外で肌移入できない
	// i4 = "Hello Go4"
	// fmt.Println(i4)

	fmt.Println(i5)

	outer()

	// Goでは未使用の変数を放置することはできない
	var s5 string = "Hello Go5"
	fmt.Println(s5)
}
```

# 数値型(Int)

```main.go
package main

import "fmt"


func main() {
	var i int = 100

	// int型全種類
	/*
	int8 int16  int32  int64
	*/

	var i2 int64 = 200

	fmt.Println(i + 50)

	// fmt.Println(i2 + i)

	// %は書式指定子で値の型を表示する
	fmt.Printf("%T\n", i2)

	// 型変換する
	fmt.Printf("%T\n", int32(i2))

	// 型変換して実行できる
	fmt.Println(i + int(i2))

}
```

# 浮動小数型(float)

```main.go
package main

import "fmt"


func main() {
	// 浮動小数点数
	var fl64 float64 = 2.4

	fmt.Println(fl64)

	fl := 3.2
	fmt.Println(fl64 + fl)

	fmt.Printf("%T, %T\n", fl64, fl)
	// float64, float64

	// float32はあんまり使わない
	var fl32 float32 = 1.2
	fmt.Printf("%T\n",fl32)
	// float32

	fmt.Printf("%T\n", float64(32))
	// float64


	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)
	// 正の無限大
	// +Inf

	ninf := -1.0 / zero
	fmt.Println(ninf)
	// 負の無限大
	// -Inf

	nan := zero / zero
	fmt.Println(nan)
	// NaN
}
```

# +の整数型(unit)

```
var u8 uint = 255   //uint型
```

# 複素数型(complex)

```
var c64 complex64 = -5 + 12i //complex型
```

# 論理型(boolean)

```main.go
package main

import "fmt"

func main() {
	var t,f bool = true,false
	fmt.Println(t,f)
}
```

# 文字列型(string)

```main.go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	fmt.Println(`test
	test
		test`)

	fmt.Println("\"")
	fmt.Println(`"`)

	fmt.Println(s[0])
	// 72
	fmt.Println(string(s[0]))
	// H

	// siをint型に変換
	i, err := strconv.Atoi(si)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
	fmt.Printf("%T\n", i)
}
```

# バイト型(byte)

```main.go
package main

import (
	"fmt"
)

func main() {
	// byte型
	// {}で囲まれた部分をスライスと呼ぶ
	byteA := []byte{72,73}
	fmt.Println(byteA)
	// [72 73]

  //アスキーコードから変換する
	fmt.Println(string(byteA))
	// HI

	c := []byte("HI")
	fmt.Println(c)
	// [72 73]

	fmt.Println(string(c))
}
```

# 配列型([])

Go の配列型は要素数を変更できない。
スライス型を使って要素数を変更することができる。

配列型　＝　要素数を変更できない。

スライス型　＝　要素数を変更可能。

```main.go
package main

import (
	"fmt"
)

// 配列型
// 他言語と異なり配列の要素数を変更できない。

func main() {
  var arr1 [3]int
	fmt.Println(arr1)
	/// [0 0 0]
	fmt.Printf("%T\n", arr1)
	// [3]int

	var arr2 [4]string = [4]string{"a", "b", "c"}
	fmt.Println(arr2)
	// [a b c ]

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)
	// [1 2 3]

	// 要素数の指定省略
	arr4 := [...]string{"a", "b", "c"}
	fmt.Println(arr4)
	// [a b c]
	fmt.Printf("%T\n", arr4)
	// [3]string

	fmt.Println(arr2[0])
	// a

	fmt.Println(arr2[2-1])
	// b

	arr2[3] = "d"
	fmt.Println(arr2)
	// [a b c d]

	// var arr5 [4]
	// arr5 = arr1
	// fmt.Println(arr5)
	// syntax error: unexpected newline, expecting type

	fmt.Println(len(arr1))
	// 3

}
```

# インターフェース型(interface)

全ての型との互換性を持っている

```main.go
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
```

# 型変換

```main.go
package main

import (
	"fmt"
	"strconv"
)

// 型変換

func main() {
	/*
	var i int = 1
	fl64 := float64(i)
	fmt.Println(fl64)
	fmt.Printf("i = %T\n", i)
	fmt.Printf("fl64 = %T\n", fl64)

	i2 := int(fl64)
	fmt.Printf("i2 = %T\n", i2)

	fl := 10.5
	i3 := int(fl)
	fmt.Printf("i3 = %T\n", i3)
	fmt.Println(i3)
	*/


	var s string = "100"
	fmt.Printf("s = %T\n", s)
	// strconv.Atoiで文字列からint型に変換
	// 2つ返り値があり、1つ目は変換に成功したかどうか、2つ目はエラーしたかどうか
	// _にすることで省略する
	i,_ := strconv.Atoi("A")


	/*
	if err != nil {
		fmt.Println(err)
		// strconv.Atoi: parsing "A": invalid syntax
	}
	*/


	fmt.Println(i)
	fmt.Printf("i = %T\n", i)

	var i2 int = 200
	s2 := strconv.Itoa(i2)
	fmt.Println(s2)
	fmt.Printf("s2 = %T\n", s2)


	// byte型変換
	var h string = "Hello world"
	b := []byte(h)
	fmt.Println(b)
	// [72 101 108 108 111 32 119 111 114 108 100]

	h2 := string(b)
	fmt.Println(h2)
	// Hello world
}
```

# 定数(グローバル変数)

```main.go
package main

import (
	"fmt"
)

// 定数(グローバルで書くことが多い)
// 頭文字を大文字にすることでグローバル変数として設定してpackageとして呼び出せる


const Pi = 3.14

const (
	URL = "http://example.com"
	SiteName = "Example"
)

const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

const (
	c0 = iota
	c1
	c2
)

// int型の最大値は「9223372036854775807」
// var Big int =	9223372036854775807 + 1
const Big = 9223372036854775807 + 1

func main() {
	fmt.Println(Pi)
	/*
		Pi = 3
	fmt.Println(Pi)
	*/

	fmt.Println(URL, SiteName)
	// http://example.com Example

	// 値を省略して定義することもできる
	fmt.Println(A, B, C, D, E, F)
	// 1 1 1 A A A

	fmt.Println(Big - 1)
	// var: constant 9223372036854775808 overflows int
	// const: 9223372036854775807

	// iotaは連続する値を自動的に返す
	fmt.Println(c0, c1, c2)
	// 0 1 2
}
```

# 算術演算子

```main.go
package main

import (
	"fmt"
)

// 算術演算子

func main() {
	fmt.Println(1 + 2)
	// 3
	fmt.Println(5 - 1)
	// 4

	fmt.Println(2 * 3)
	// 6

	fmt.Println(60 / 3)
	// 20

	fmt.Println(9 % 4)
	// 1

	fmt.Println("ABC" + "DEF")
	// ABCDEF

	n := 0
	n += 4 // n = n + 4
	fmt.Println(n)
	// 4
	n++ // n = n + 1
	fmt.Println(n)
	// 5

	n-- // n = n - 1
	fmt.Println(n)
	// 4

	s := "ABC"
	s += "DEF" // s = s + "DEF"
	fmt.Println(s)
	// ABCDEF
}
```

# 比較演算子

```main.go
package main

import (
	"fmt"
)

// 比較演算子

func main() {
	fmt.Println(1 == 1)
	// true

	fmt.Println(1 == 2)
	// false

	fmt.Println(4 <= 8)
	// true

	fmt.Println(4 >= 8)
	// false

	fmt.Println(4 < 8)
	// true

	fmt.Println(true == false)
	// false

	fmt.Println(true != false)
	// true
}
```

# 論理演算子

```main.go
package main

import (
	"fmt"
)

// 論理演算子

func main() {
	fmt.Println(true && false == true)
	// false

	fmt.Println(true && true == true)
	// true

	// ||はどちらかがteureであればtrue
	fmt.Println(true || false == true)
	// true

	fmt.Println(false || false == true)
	// false

	fmt.Println(!true)
	// false

	fmt.Println(!false)
	// true

	fmt.Println(!!true)
	// true

	fmt.Println(!!false)
	// false
}
```

# 関数

```main.go
package main

import (
	"fmt"
)

// 関数

// 第一引数にint型のx、第二引数にint型yを受け取りint型を返り値
// 引数が同じ型なら、型推論で型推論できる
func Plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q,r
}

// resultを関数内で使ってreturnしているので省略できる
func Duble(price int) (result int) {
	result = price * 2
	return
}

func Noreturn() {
	fmt.Println("no return")
	return
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i)

	i2, i3 := Div(9, 3)
	fmt.Println(i2,i3)
	// 3 0

  i4, _ := Div(9, 3)
	fmt.Println(i4)
	// 3

	i5 := Duble(1000)
	fmt.Println(i5)
	// 2000

	Noreturn()
	// no return
}
```

# 無名関数

```main.go
package main

import (
	"fmt"
)

// 関数
// 無名関数


func main() {
	f := func(x,y int) int {
		return x + y
	}
	i := f(1, 2)
	fmt.Println(i)
	// 3

	// 即時で引数を入れることもできる
	i2 := func(x, y int) int {
		return x + y
	}(1, 2)
	fmt.Println(i2)
	// 3
}
```

# 関数を返す関数

```main.go
package main

import (
	"fmt"
)

// 関数
// 関数を返す関数

func ReturnFunc() func() {
	return func() {
		fmt.Println("i am a function")
	}
}

func main() {
	f := ReturnFunc()
	f()
	// i am a function
}
```

# 関数を引数に取る関数

```main.go
package main

import (
	"fmt"
)

// 関数
// 関数を引数に取る関数

func CallFunc(f func()) {
	f()
}

func main() {
	CallFunc(func() {
		fmt.Println("I am a Function")
	})
}
```

# クロージャー

Go の無名関数はクロージャーで、

クロージャーとは日本語では関数閉包と呼ばれ、関数と関数の処理に関する関数外の環境をセットして閉じ込めた物です。

return されても store は初期化されないところが重要。

```main.go
package main

import "fmt"

// クロージャー

func Later() func(string) string {
	// storeは空文字でる
	// 2度目の実行でfunc(next string)でnextに代入されていた値が入ってくる
	var store string
	// return func(next string)から始まるので"Hello"を代入した段階ではまだstoreは空文字
	// storeはreturn func(next string)が実行され続ける限り監視され続ける
	// returnされてもstoreは初期化されないところが重要
		return func(next string) string {
		// 最初はstoreは空文字なので""が出力される
		s := store
		// nextで渡された文字列"Hello"が代入されている
		store = next
		// 返り値はstore("空文字")なので最初は空文字が返る
		return s
	}
}

func main() {
	f := Later()
	fmt.Println(f("Hello"))
	fmt.Println(f("My"))
	fmt.Println(f("name"))
	fmt.Println(f("is"))
	fmt.Println(f("Golang"))
	// ""
	// "Hello"
	// "My"
	// "name"
	// "is"
}
```

# ジェネレーター

何らかのルールに従って連続した値を返し続ける仕組みの事。

```main.go
package main

import "fmt"

// ジェネレーター
// クロージャーを応用する

func integers() func() int {
	i := 0
	// iに対して増分して返す関数
	return func() int {
		i++
		return i
	}
}

func main() {
	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	// 1
	// 2
	// 3
	// 4

	// 再定義してもintsと同じ挙動をする
	others := integers()
	fmt.Println(others())
	fmt.Println(others())
	fmt.Println(others())
	fmt.Println(others())
	// 1
	// 2
	// 3
	// 4
}
```

# if(条件分岐)

```main.go
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
```

# エラーハンドリング

```main.go
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
```

# for(繰り返し処理)

```main.go
package main

import "fmt"

// for
// 繰り返し処理

func main() {
		i := 0
	for {
		i++
		// forを中断させる場合は"break"を使う
		if i > 10 {
			break
		}
		fmt.Println("Loop")
	}

		// 条件付き繰り返し処理
	point := 0
	for point < 10 {
		fmt.Println(point)
		point++
	}
	// 0
	// 1
	// 2
	// 3
	// ...
	// 9

	for i := 0; i < 10; i++ {
		if i == 3 {
			// continueすることでiが3の場合は処理をスキップする
			continue
		}
		fmt.Println(i)
			// 0
			// 1
			// 2
			// 4
			// ...
			// 9
	}

	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
		// 1
		// 2
		// 3
	}

	arr2 := [3]int{1, 2, 3}
	// iがindex番号でvが値
	// 要略したい場合は_を使う
	for i, v := range arr2 {
		fmt.Println(i, v)
		// 0 1
		// 1 2
		// 2 3
	}

	// スライス
	// 可変長
	sl := []string{"Python", "PHP", "Go"}
	for i, v := range sl {
		fmt.Println(i, v)
		// 0 Python
		// 1 PHP
		// 2 Go
	}

	// マップ
	// keyとvalueを取り出す
	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
		// apple 100
		// banana 200
	}
}
```

# switch

```main.go
package main

import "fmt"

// switch
// 式スイッチ

func main() {
	n := 1
	switch n {
	case 1, 2:
		fmt.Println("one or two")
	case 3, 4:
		fmt.Println("three or four")
	default:
		fmt.Println("other")
	}

	switch n2 := 2; n2 {
	case 1, 2:
		fmt.Println("one or two")
	case 3, 4:
		fmt.Println("three or four")
	default:
		fmt.Println("other")
	}

	n3 := 6
	switch {
	case n3 > 0 && n3 < 4:
		fmt.Println("0 < n3 < 4")
	case n3 > 3 && n3 < 7:
		fmt.Println("3 < n3 < 7")
	}

		switch n4 := 2; n4 {
	case 1, 2:
		fmt.Println("one or two")
	case 3, 4:
		fmt.Println("three or four")
		// 列挙と判定が同じswitchに混在する場合はエラーになる
		// case n4 > 3 && n4 < 6:
		// fmt.Println("3 < n4 < 6")
	default:
		fmt.Println("other")
	}
}
```

# 型 Switch&型アサーション

```main.go
package main

import "fmt"

// switch
// 型スイッチ

// 型アサーション
func anything(a interface{}) {
	switch v := a.(type) {
	case int:
		fmt.Println(v, "is int")
	case string:
		fmt.Println(v, "is string")
	default:
		fmt.Println(v, "is unknown")
	}
}

func main() {
	anything("aaa")
	anything(1)

	var x interface{} = 3
	// interfaceをintで復元する
	i := x.(int)
	fmt.Println(i + 2)
	// 5

	// xはinterfaceで型が外れる
	// fmt.Println(x + 2)
	// error


	// int型なのでfloat64に変換できない
	// f := x.(float64)
	// fmt.Println(f + 2)
	// runtime error

	// runtime errorにならない
	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64)
	// 0 false

	i2, isInt := x.(int)
	fmt.Println(i2, isInt)
	// 3 true

	if x == nil {
		fmt.Println("x is nil")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i,"x is int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s,"x is string")
	} else {
		fmt.Println("x is unknown")
	}

	switch x.(type) {
	case int:
		fmt.Println("x is int")
	case string:
		fmt.Println("x is string")
	default:
		fmt.Println("x is unknown")
	}

	switch v := x.(type) {
	case bool:
		fmt.Println(v, "is bool")
	case int:
		fmt.Println(v, "is int")
	case string:
		fmt.Println(v, "is string")
	default:
		fmt.Println(v, "is unknown")
	}
}
```

# ラベル付き for

```main.go
package main

import "fmt"

// ラベル付きfor

func main() {
	/*
	Loop:
	for {
		for {
			for {
				fmt.Println("START")
				// 上で宣言したLoopをラベルすることで間の処理をスキップすることができる
				break Loop
			}
			fmt.Println("処理をしない")
		}
		fmt.Println("処理をしない")
	}
	fmt.Println("END")
	*/

	// jが1より上であれば「fmt.Println("処理をしない")」の処理をスキップしてjのforを繰り返す
	Loop:
	for i := 0; i < 3; i++ {
		for j := 1; j < 3; j++ {
			if j > 1 {
				continue Loop
			}
			fmt.Println(i,j,i+j)
		}
		fmt.Println("処理をしない")
	}

}
```

# defer

関数の最後に実行される

```main.go
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
```

# Panic & Recover

```main.go
package main

import "fmt"

// panic & recover
// 例外処理(runtimeを強制的に停止させる)
// あまり使わない

func main() {
	defer func() {
		// panicが発生した時に実行されるようにするためdeferでrecoverを定義する方がいい
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	panic("runtime errpr")
	fmt.Println("one or two")
}
```

# Goroutin(ゴルーチン)並行処理

```main.go
package main

import (
	"fmt"
	"time"
)

// go goroutin(ゴルーチン)
// 並行処理

func sub() {
	for {
		fmt.Println("sub loop")
		// 100ms感覚で実行される
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// 「go」を付けることで並行処理できる
	// 「go」がなければsub()が無限ループして永遠に下のforに到達しない
	go sub()
	go sub()

	for {
		fmt.Println("main loop")
		time.Sleep(200 * time.Millisecond)
	}
}
```

# init(初期化)

```main.go
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
```

# スライス

```main.go
package main

import (
	"fmt"
)

// スライス
// 宣言・操作

func main() {
	// 配列は[]内に要素数を指定していた
	var sl []int
	fmt.Println(sl)
	// []

	var sl2 []int = []int{100,200}
	fmt.Println(sl2)
	// [100 200]

	sl3 := []string{"A", "B"}
	fmt.Println(sl3)
	// [A B]

	sl4 := make([]int, 5)
	fmt.Println(sl4)
	// [0 0 0 0 0]

	sl2[0] = 1000
	fmt.Println(sl2)
	// [1000 200]

	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5)

	fmt.Println(sl5[0])

	// indexの2と4
	fmt.Println(sl5[2:4])
	// [3 4]

	// indexの2の手前まで
	fmt.Println(sl5[:2])
	// [1 2]

	// indexの3から最後
	fmt.Println(sl5[2:])
	// [3 4 5]

	// indexの0から最後
	fmt.Println(sl5[:])
	// [1 2 3 4 5]

	// 最後の要素
	fmt.Println(sl5[len(sl5)-1])

	// 一番最初と最後以外の要素
	fmt.Println(sl5[1 : len(sl5)-1])
	// [2 3 4]
}
```

# スライス(make / append / cap)

make でスライスを作る。第二引数で容量を確保する
cap で容量を測る
append で要素の追加を行う

要領以上の要素が追加されるとメモリの消費が倍になってしまいます。

メモリーを気にするような開発の場合は、容量にも気をつけます。

最初は気にせずやるほうがいいと思います。

過剰にメモリを確保してしまうと実行速度が落ちたりする。

良質なパフォーマンスを実現するには、容量の管理も気にします

```main.go
package main

import (
	"fmt"
)

// スライス
// append make len cap

func main() {
	sl := []int{100, 200}
	fmt.Println(sl)

	// 300という要素を追加する
	sl = append(sl, 300)
	fmt.Println(sl)
	// [100 200 300]

	sl = append(sl, 400, 500)
	fmt.Println(sl)
	// [100 200 300 400 500]

	sl2 := make([]int, 5)
	fmt.Println(sl2)

	fmt.Println(len(sl2))
	// 5

	// sl2の容量を測る
	fmt.Println(cap(sl2))
	// 5

	// 第二引数で容量を指定する
	sl3 := make([]int, 5, 10)

	fmt.Println(sl3)
	// [0 0 0 0 0]

	fmt.Println(len(sl3))
	// 5

	fmt.Println(cap(sl3))
	// 10

	sl3 = append(sl3, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(sl3)
	// [0 0 0 0 0 1 2 3 4 5 6 7]

	fmt.Println(cap(sl3))
	// 20

}
```

# スライス(copy)

```main.go
package main

import (
	"fmt"
)

// スライス
// copy

func main() {
	sl := []int{100, 200}
	sl2 := sl

	// 元データのslに1000が入る
	sl2[0] = 1000
	fmt.Println(sl)
	// [1000 200]

	// 基本型の場合はそれぞれ独立している
	// 山椒型の場合は元データも更新される
	var i int = 10
	i2 := i
	i2 = 100
	fmt.Println(i)
	// 10
	fmt.Println(i2)
	// 100

	sl3 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl3)
	sl4 := make([]int, 5, 10)
	fmt.Println(sl4)
	// 第一引数にコピー先、第二引数にコピー元
	// 先頭から塗りつぶすようにコピーしていくしよう
	n := copy(sl4 ,sl3)
	fmt.Println(n, sl4)
	// 5 [1 2 3 4 5]
}
```

# スライス(for)

```main.go
package main

import (
	"fmt"
)

// スライス
// for

func main() {
	sl := []string{"A", "B", "C"}
	fmt.Println(sl)
	// [A B C]

	for i, v := range sl {
		// index番号と要素を取得
		fmt.Println(i, v)
		// 0 A
		// 1 B
		// 2 C
	}

	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
		// A
		// B
		// C
	}

}
```

# スライス(可変長引数)

```main.go
package main

import (
	"fmt"
)

// スライス
// 可変長引数

func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	fmt.Println(Sum(1, 2, 3, 4, 5))
	// 15

	fmt.Println(Sum())
	// 0

	sl := []int{1, 2, 3}
	fmt.Println(Sum(sl...))
	// 6
}
```

# マップ

```main.go
package main

import (
	"fmt"
)

// マップ

func main() {
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m)
	// map[A:100 B:200]

	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2)
	// map[A:100 B:200]

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)
	// map[1:A 2:B]

	m4 := make(map[int]string)
	fmt.Println(m4)
	// map[]

	m4[1] = "JAPAN"
	m4[2] = "USA"
	fmt.Println(m4)
	// map[1:JAPAN 2:USA]

	fmt.Println(m4[1])
	// JAPAN
	fmt.Println(m["A"])
	// 100
	fmt.Println(m4[3])
	// "" <-nilにもならない

	s, ok := m4[1]
	fmt.Println(s, ok)
	// JAPAN true

	s, ok = m4[3]
	fmt.Println(s, ok)
	//  "" false
	if !ok {
		fmt.Println("no data")
	}

	m4[2] = "US"
	fmt.Println(m4)
	// map[1:JAPAN 2:US]

	m4[3] = "CHINA"
	fmt.Println(m4)
	// map[1:JAPAN 2:US 3:CHINA]

	// deleteの第一引数でmapを第二引数にkeyを指定して削除する
	delete(m4, 3)
	fmt.Println(m4)
	// map[1:JAPAN 2:US]

}
```

# マップ(for)

```mian.go
package main

import (
	"fmt"
)

// マップ
// for

func main() {
	m := map[string]int{
		"Apple": 100,
		"Banana": 200,
	}

	for k, v := range m {
		fmt.Println(k, v)
		// Apple 100
		// Banana 200
	}
}
```

# チャネル

ゴルーチン間でデータの送受信を行う

```main.go
package main

import (
	"fmt"
)

// channel
// 複数のゴルーチン間でのデータの受け渡しをするために設計されたデータ構造
// 宣言、操作

func main() {
	// チャネルの宣言
	var ch1 chan int

	/*
	受信用のチャネルとして明示的に定義
	var ch2 <-chan int

	送信用のチャネルとして明示的に定義
	var ch3 chan<- int
	*/

	// makeを使ってチャネルに書き込みや読み込みを行えるようにする
	ch1 = make(chan int)

	ch2 := make(chan int)

	fmt.Println(cap(ch1))
	// 0
	fmt.Println(cap(ch2))
	// 0

	// bufサイズを第二引数で指定する
	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3))
	// 5

	// ch3に1を送信する
	ch3 <- 1
	fmt.Println(len(ch3))
	// 1

	ch3 <- 2
	ch3 <- 3
	fmt.Println("len", len(ch3))
	// 3

	i := <-ch3
	fmt.Println(i)
	// 1
	fmt.Println("len", len(ch3))
	// 2

	i2 := <-ch3
	fmt.Println(i2)
	// 2
	fmt.Println("len", len(ch3))
	// 1

	fmt.Println(<-ch3)
	// 3
	fmt.Println("len", len(ch3))
	// 0

	ch3 <- 1
	// ここで一つ取り出すことでdeadlockにはならない
	fmt.Println(<-ch3)
	// 1
	ch3 <- 2
	fmt.Println(<-ch3)
	// 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	ch3 <- 6
	// deadlock!
	// bufを超えた場合はdeadlockエラーになる

	// 先に入れたデータから順番に取り出される
}
```

# チャネル(channel / goroutin)

```main.go
package main

import (
	"fmt"
	"time"
)

// channel
// チャネルとゴルーチン

func reciever(c chan int) {
	for {
		i := <- c
		fmt.Println(i)
	}
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	// deadlock
	// ゴルーチン間での送受信が前提なのでチャネル単体ではdeadlockになる

	// チャネルに値が入るのを待つ
	go reciever(ch1)
	go reciever(ch2)

	// i(int)をforでチャネルに送信
	i := 0
	for i < 100 {
		ch1 <- i
		ch2 <- i
		time.Sleep(50 * time.Millisecond)
		i++
	}
	// 1
	// 1
	//1
	// 2
	// ...
	// 99
	// 99
}
```

# チャネル(close)

```main.go
package main

import (
	"fmt"
	"time"
)

// Channel
// close

func recieve(name string, ch chan int) {
	for {
		// 受け取ったchがclose状態かどうかをチェック
		i, ok := <- ch
		// もしclose状態ならループを抜ける
		if !ok {
			break
		}
		// closeでなければ受け取った値を表示
		fmt.Println(name, i)
	}
	fmt.Println(name, "end")
}

func main() {
	// 最初はopen状態
	ch1 := make(chan int, 2)


	// ch1 <- 1

	// close(ch1)

	// ch1 <- 1
	// runtimeError

	// fmt.Println(ch1)
	// 0xc0000b6000

	// runtimeエラーにならない
	// ch1がopen状態であればtrueになる
	// i, ok := <- ch1
	// fmt.Println(i, ok)
	// 0 false

	// i2, ok := <- ch1
	// fmt.Println(i2, ok)
	// 0 false

	go recieve("1.goroutin", ch1)
	go recieve("2.goroutin", ch1)
	go recieve("3.goroutin", ch1)

	i := 0
	for i < 100 {
		ch1 <- i
		i++
	}

	close(ch1)
	time.Sleep(3 + time.Second)


}
```

# チャネル(for)

```main.go
package main

import (
	"fmt"
)

// Channel
// for

func main() {
	ch1 := make(chan int, 3)

	ch1 <- 1
	ch1 <- 2
	ch1 <- 3

	// closeで締めないとdeadrockになる
	close(ch1)
	for i := range ch1 {
		fmt.Println(i)
	}
	/*
	1
	2
	3
	fatal error: all goroutines are asleep - deadlock!
	*/

}
```

# チャネル(select)

```main.go
package main

import "fmt"

// Channel
// select

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)

	// これらがあるとランダムに出力されてしまう
	ch2 <- "A"
	ch1 <- 1
	ch2 <- "B"
	ch1 <- 2

	// どちらかのチャネルをselectして行えばデッドロックされない
	// しかし、上のv1のみ計算される。ランダムになってしまう
	select {
	case v1 := <- ch1:
		fmt.Println(v1 + 1000)
	case v2 := <- ch2:
		fmt.Println(v2 + "!?")
	default:
		fmt.Println("どちらでもない")
	}

	ch3 := make(chan int)
	ch4 := make(chan int)
	ch5 := make(chan int)

	// reciever
	go func() {
		for {
			i := <- ch3
			ch4 <- i * 2
		}
	}()

	go func() {
		for {
			i2 := <- ch4
			ch5 <- i2 - 1
		}
	}()

	n := 0
	// Lを定義しないとbreakしても無限ループ
	L:
	for {
		select {
			case ch3 <- n:
				n++
			case i3 := <- ch5:
				fmt.Println("recieve", i3)
			default:
				if n > 100 {
					break L
				}
		}
		// if n > 100 {
		// 	break
		// }
	}
	/*
	recieve -1
	recieve 0
	recieve 1
	...
	recieve 197
	*/
}
```

# ポインタ型

![](https://img-c.udemycdn.com/redactor/raw/article_lecture/2021-07-08_17-17-13-dafb07c608a3a20380d54cf2b7249d6a.PNG)

Go にはポインタという機能があります。

ポインタとは、値型に分類されるデータ構造(基本型、参照型、構造体）のメモリ上のアドレスと型の情報のことです。

Go ではこれを使って、データ構造を間接的に参照、操作ができる様になっている。

主に値型の操作に使われ、参照型は元からその機能を含んでいる為、基本的には不要となっています。

> var などで定義したデータを更新するときにポインタを使って、対象のデータがどこのアドレスに格納されているかを確認する。
> 対象のデータが格納されているアドレスに対して値の更新を行うことでデータを更新することができる
> 参照型であればポインタを使わずに値の更新ができる

```main.go
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
```

# struct(構造体)

クラス的なもの

```main.go
package main

import "fmt"

// struct
// クラス的なもの

type User struct {
	Name string
	Age int
	// X, Y int
}

func UpdateUser(user User) {
	user.Name = "user1"
	user.Age = 1000
}

func UpdateUser2(user *User) {
	user.Name = "A"
	user.Age = 1000
}


func main() {
	var user1 User
	fmt.Println(user1)
		// { 0 0 0}
	user1.Name = "user1"
	user1.Age = 20
	fmt.Println(user1)
	// {user1 20 0 0}

	user2 := User{}
	fmt.Println(user2)
	// { 0 0 0}

	user2.Name = "user2"
	user2.Age = 30
	fmt.Println(user2)
	// {user2 30 0 0}

	user3 := User{
		Name: "user3",
		Age: 40,
	}
	fmt.Println(user3)
	// {user3 40 0 0}

	// X, Yをなくすとkeyを指定しなくてもできる
	user4 := User{"user4", 50}
	fmt.Println(user4)
	// {user4 50}

	user6 := User{Name: "user6"}
	fmt.Println(user6)
	// {user6 0}

	// newで定義するとポインタ型になる
	user7 := new(User)
	fmt.Println(user7)
	// &{ 0}

	// アドレス演算子でポインタを取得
	// ポインタを使うときはアドレス演算子を使う
	user8 := &User{}
	fmt.Println(user8)
	// &{ 0}

	// 値型なので本体であるuser1には影響が出ない
	UpdateUser(user1)
	// ポインタ型で渡してあげると更新できる
	UpdateUser2(user8)
	fmt.Println(user1)
	// {user1 20}
	fmt.Println(user8)
	// &{A 1000}

}
```

# struct(メソッド)

メソッドの引数は基本的にポインタ型にしておくのが吉

```main.go
package main

import "fmt"

// struct
// メソッド

type User struct {
	Name string
	Age int
	// X, Y int
}

// メソッドの宣言
// メソッドはfuncの後ろに()でくくったものがメソッド名になる
func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u User) SetName(name string) {
	u.Name = name
}

func (u *User) SetName2(name string) {
	u.Name = name
}

func main() {
	user1 := User{
		Name: "user1",
	}
	user1.SayName()
	// user1

	user1.SetName("user2")
	user1.SayName()
	// user1

	// SetName2はポインタ型なので本体のNameを更新する
	user1.SetName2("user3")
	user1.SayName()
	// user3

	user2 := &User{
		Name: "user2",
	}
	user2.SetName2("B")
	user2.SayName()
	// B
}
```

# struct(埋め込み)

```main.go
package main

import "fmt"

// struct
// 埋め込み

// Tの構造体にUserを埋め込むことができる
type T struct {
	User
}

type User struct {
	Name string
	Age int
	// X, Y int
}

func (u *User) SetName() {
	u.Name = "A"
}

func main() {
	t := T{User: User{Name: "user1", Age: 20}}
	fmt.Println(t)
	// {{user1 20}}

	fmt.Println(t.User)
	// {user1 20}

	// TにUserのみで型名を省略した場合のみアクセス可能
	fmt.Println(t.Name)
	// user1

	t.SetName()
	fmt.Println(t)
	// {{A 20}}
}
```

# struct(コントラクト)

```mainb.go
package main

import "fmt"

// struct
// コントラクト

type User struct {
	Name string
	Age int
	// X, Y int
}

func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age: age,
	}
}

func main() {
	user1 := NewUser("user1", 20)
	fmt.Println(user1)
	// &{user1 20}

	// 実態にアクセスする
	fmt.Println(*user1)
	// {user1 20}
}
```

# struct(スライス)

```main.go
package main

import "fmt"

// struct
// コントラクト

type User struct {
	Name string
	Age int
	// X, Y int
}

type Users []*User

/*
type Users struct {
	Users []*User
}
*/


func main() {
	user1 := User{
		Name: "user1",
		Age: 10,
	}

	user2 := User{
		Name: "user2",
		Age: 20,
	}

	user3 := User{
		Name: "user3",
		Age: 30,
	}

	user4 := User{
		Name: "user4",
		Age: 40,
	}

	users := Users{}
	users = append(users, &user1, &user2, &user3, &user4)
	fmt.Println(users)
	// [0xc00000c030 ...]

	for _, u := range users {
		fmt.Println(*u)
	}
	// {user1 10}
	// {user2 20}
	// {user3 30}
	// {user4 40}

	users2 := make([]*User, 0)
	users2 = append(users2, &user1, &user2, &user3, &user4)

	for _, u := range users2 {
		fmt.Println(*u)
	// {user1 10}
	// {user2 20}
	// {user3 30}
	// {user4 40}
	}
}
```

# struct(マップ)

```main.go
package main

import "fmt"

// struct
// map

type User struct {
	Name string
	Age int
	// X, Y int
}


func main() {
	m := map[int]User{
		1: {
			Name: "user1",
			Age: 20,
		},
		2: {
			Name: "user2",
			Age: 30,
		},

	}
	fmt.Println(m)
	// map[1:{user1 20} 2:{user2 30}]

	m2 := map[User]string{
		{Name: "user1", Age: 20}: "Tokyo",
		{Name: "user2", Age: 30}: "Osaka",
	}
	fmt.Println(m2)
	// map[{user1 20}:Tokyo {user2 30}:Osaka]

	m3 := make(map[int]User)
	fmt.Println(m3)
	// map[]

	m3[1] = User{Name: "user3"}
	fmt.Println(m3)
	// map[1:{user3}]

	for _, v := range m {
		fmt.Println(v)
		// {user1 20}
		// {user2 30}
	}

}
```

# struct(独自型)

```main.go
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
```

# interface(型アサーション)

```main.go
package main

import "fmt"

// interface
// もっともポピュラーな使い方。異なる方に共通の性質を付与する

type Stringfy interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v, Model=%v", c.Number, c.Model)
}

func main() {
	vs := []Stringfy{
		&Person{Name: "Taro", Age: 21},
		&Car{Number: "123-456", Model: "AB-1234"},
	}

	// interfaceを使うことで異なる構造体に共通の処理やメソッドを実装することができる
	for _, v := range vs {
		fmt.Println(v.ToString())
		// Name=Taro, Age=21
		// Number=123-456, Model=AB-1234
	}
}
```

# interface(カスタムエラー)

```main.go
package main

import "fmt"

// interface
// カスタムエラー

/*
Golangのソースコード
type error interface {
	Error() string
}
*/

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{
		Message: "This is my error",
		ErrCode: 1234,
	}
}

func main() {
	err := RaiseError()
	fmt.Println(err.Error())

	// fmt.Println(err.Message)
	// error

	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.ErrCode)
		// This is my error
	}

}
```

# interface(Stringer)

```main.go
package main

import "fmt"

// interface
// fmt.Stringer

/*
type Stringer interface {
	String() string
}
*/

type Point struct {
	A int
	B string
}

func (p *Point) String() string {
	return fmt.Sprintf("<<%v, %v>>", p.A, p.B)
}

func main() {
	p := &Point{100, "ABC"}
	fmt.Println(p)
	// <<100, ABC>>

}
```

# モジュール管理の補足

次のレクチャーでの、パッケージを import して実行した際に

is not in GOROOT

表題のエラーが出てしまい、他のフォルダから変数の参照ができない場合。

ターミナルで実行してみてください。

go mod init

go mod tidy

下記記事が参考になるかと思います。

https://qiita.com/taku-yamamoto22/items/4d6f9ff8451a0b86997b
