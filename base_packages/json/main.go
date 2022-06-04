package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// json

type A struct {

}

type User struct {
	// stringを指定すると文字列にもなる
	ID int `json:"id,string"`
	// 空文字であればomitemptyを指定するとJSONに含まれない
	Name string `json:"name,omitempty"`
	Email string `json:"email"`
		// -を指定すると隠せる
	Created time.Time `json:"-"`
	// アドレス型を指定してomitemptyを指定するとnullに含まれない
	A *A `json:"omitempty"`
}

// "MarshalJSON"という名前で眼メソッドを定義している
func (u User) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr " + u.Name,
	})
	return v, err
}

// "UnmarshalJSON"という名前で眼メソッドを定義している
// ポインタ型の変数に対してUnmarshalJSONを呼び出すとエラーになる
func (u *User) UnmarshalJSON(b []byte) error {
	type User2 struct {
		Name string
	}
	var u2 User2
	err := json.Unmarshal(b, &u2)
	if err != nil {
		return err
	}
	u.Name = u2.Name + "!"
	return err
}

func main() {
	u := new(User)
	u.ID = 1
	u.Name = "test"
	u.Email = "example@example.com"
	u.Created = time.Now()

	// MarchalJSONに変換
	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}
	// byte型で出てくるので文字列に変換
	fmt.Println(string(bs))

	fmt.Printf("%T\n", bs)
	u2 := new(User)

	// UnmarchalでJSONから構造体に変換する
	if err := json.Unmarshal(bs, &u2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(u2)
	
}