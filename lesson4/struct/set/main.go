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