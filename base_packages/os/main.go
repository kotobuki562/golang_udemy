package main

import (
	"fmt"
	"log"
	"os"
)

// os

func main() {
	// Exit
	// Goはmainが割れば終了する
	// Exitを使えば任意のタイミングで終了することができる
	// 引数で終了ステータスを指定できる
	/*
	os.Exit(1)
	fmt.Println("Start")
	*/

	// ファイル操作
	f, _ := os.Create("test.txt")
	f.Write([]byte("Hello\n"))
	f.WriteAt([]byte("Golang"),6)
	f.Seek(0, os.SEEK_END)
	f.WriteString("Yaah")

	// ファイル操作
	// test.txtを開く
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()
	fmt.Println(f)


	// log
	// 	_,err := os.Open("A.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// 下記がビルドして実行するコマンド
	// go build -o getargs main.go
	// ./getargsができる
	//  ./getargs -a=1 -b=2 -c=3で引数を指定して実行できる
	// Args
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	fmt.Println(os.Args[3])

	// os.Argsの要素数を表示
	fmt.Printf("length=%d\n", len(os.Args))
	// os.Argsの中身を表示
	for i, v := range os.Args {
		fmt.Printf("%d:%s\n", i, v)
	}

		// deferも実行されない
	// defer func() {
	// 	fmt.Println("defer")
	// }()
	// os.Exit(0)

}