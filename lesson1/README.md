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
