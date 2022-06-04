package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// net/http
// サーバー

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func top(w http.ResponseWriter, r *http.Request) {
	// ファイルを解析
	t, err := template.ParseFiles("top.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// html内の{{ . }}に渡す値を指定
	t.Execute(w, "Hello HTML")
}



func main() {
	// handlerを分けて各ページの処理を行う
	http.HandleFunc("/top", top)
	http.ListenAndServe(":8080", nil)
	// hello, world!と表示される
}