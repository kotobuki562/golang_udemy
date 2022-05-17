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