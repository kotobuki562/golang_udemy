package main

import "fmt"

// struct
// クラス的なもの

type User struct {
	Name string
	Age int
	// X, Y int
}

func UpdateUser(user User) {
	user.Name = "user1"
	user.Age = 1000
}

func UpdateUser2(user *User) {
	user.Name = "A"
	user.Age = 1000
}


func main() {
	var user1 User
	fmt.Println(user1)
		// { 0 0 0}
	user1.Name = "user1"
	user1.Age = 20
	fmt.Println(user1)
	// {user1 20 0 0}
	
	user2 := User{}
	fmt.Println(user2)
	// { 0 0 0}

	user2.Name = "user2"
	user2.Age = 30
	fmt.Println(user2)
	// {user2 30 0 0}

	user3 := User{
		Name: "user3",
		Age: 40,
	}
	fmt.Println(user3)
	// {user3 40 0 0}

	// X, Yをなくすとkeyを指定しなくてもできる
	user4 := User{"user4", 50}
	fmt.Println(user4)
	// {user4 50}
	
	user6 := User{Name: "user6"}
	fmt.Println(user6)
	// {user6 0}

	// newで定義するとポインタ型になる
	user7 := new(User)
	fmt.Println(user7)
	// &{ 0}

	// アドレス演算子でポインタを取得
	// ポインタを使うときはアドレス演算子を使う
	user8 := &User{}
	fmt.Println(user8)
	// &{ 0}
	
	// 値型なので本体であるuser1には影響が出ない
	UpdateUser(user1)
	// ポインタ型で渡してあげると更新できる
	UpdateUser2(user8)
	fmt.Println(user1)
	// {user1 20}
	fmt.Println(user8)
	// &{A 1000}

}