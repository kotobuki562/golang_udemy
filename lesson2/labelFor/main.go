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