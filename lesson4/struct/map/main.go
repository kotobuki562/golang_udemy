package main

import "fmt"

// struct
// map

type User struct {
	Name string
	Age int
	// X, Y int
}


func main() {
	m := map[int]User{
		1: {
			Name: "user1",
			Age: 20,
		},
		2: {
			Name: "user2",
			Age: 30,
		},

	}
	fmt.Println(m)
	// map[1:{user1 20} 2:{user2 30}]

	m2 := map[User]string{
		{Name: "user1", Age: 20}: "Tokyo",
		{Name: "user2", Age: 30}: "Osaka",
	}
	fmt.Println(m2)
	// map[{user1 20}:Tokyo {user2 30}:Osaka]
	
	m3 := make(map[int]User)
	fmt.Println(m3)
	// map[]

	m3[1] = User{Name: "user3"}
	fmt.Println(m3)
	// map[1:{user3}]

	for _, v := range m {
		fmt.Println(v)
		// {user1 20}
		// {user2 30}
	}

}