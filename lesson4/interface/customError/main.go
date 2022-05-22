package main

import "fmt"

// interface
// カスタムエラー

/*
Golangのソースコード
type error interface {
	Error() string
}
*/

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{
		Message: "This is my error",
		ErrCode: 1234,
	}
}

func main() {
	err := RaiseError()
	fmt.Println(err.Error())

	// fmt.Println(err.Message)
	// error

	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.ErrCode)
		// This is my error
	}

}