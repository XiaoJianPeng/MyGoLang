package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/go", myhandler)
	http.ListenAndServe("127.0.0.1:8000", nil)
}

func myhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接成功")
	// 请求方式：GET POST DELETE PUT UPDATE
	fmt.Println("method", r.Method)

	fmt.Println("url:", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)
	// 回复
	w.Write([]byte("hello world!"))
}
