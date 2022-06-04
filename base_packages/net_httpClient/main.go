package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// cnet/url
// URLの文字列を解析する


func main() {

	// GETメソッド
	res, _ := http.Get("https://example.com")
	fmt.Println(res.StatusCode)

	fmt.Println(res.Proto)

	fmt.Println(res.Header["Date"])
	fmt.Println(res.Header["Content-type"])
	fmt.Println(res.Request.Method)
	fmt.Println(res.Request.URL)

	// リソースの解放
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))


	// POSTメソッド
	vs := url.Values{}
	vs.Add("id", "1")
	vs.Add("message", "メッセージ")
	fmt.Println(vs.Encode())
	// id=1&message=%E3%83%A1%E3%83%83%E3%82%BB%E3%83%BC%E3%82%B8

	res, err := http.PostForm("https://example.com/", vs)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	resBody, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(resBody))
	

}