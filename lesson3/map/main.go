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