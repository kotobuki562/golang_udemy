package main

// cnet/url
// URLの文字列を解析する

import (
	"fmt"
	"net/url"
)


func main() {
	// URLを解析
	u, _ := url.Parse("http://example.com/search?a=1&b=2#top")
	fmt.Println(u.Scheme)
	// http
	fmt.Println(u.Host)
	// example.com
	fmt.Println(u.Path)
	// /search
	fmt.Println(u.RawQuery)
	// a=1&b=2
	fmt.Println(u.Fragment)
	// top


	// URLを生成
	
	url := &url.URL{}
	url.Scheme = "https:"
	url.Host = "google.com"
	q := url.Query()
	q.Set("q", "Go言語")

	url.RawQuery = q.Encode()

	fmt.Println(url)
	// https:://google.com/?q=Golang
	// 日本語入力するとエンコードしてくれる
	// https:://google.com?q=Go%E8%A8%80%E8%AA%9E

}