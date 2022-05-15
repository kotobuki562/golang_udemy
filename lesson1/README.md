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
