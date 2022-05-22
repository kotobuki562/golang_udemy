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