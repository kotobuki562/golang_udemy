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