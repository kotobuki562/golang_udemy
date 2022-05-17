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