package main

import "fmt"

// struct
// コントラクト

type User struct {
	Name string
	Age int
	// X, Y int
}

func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age: age,
	}
}

func main() {
	user1 := NewUser("user1", 20)
	fmt.Println(user1)
	// &{user1 20}

	// 実態にアクセスする
	fmt.Println(*user1)
	// {user1 20}
}