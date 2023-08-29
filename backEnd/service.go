package main

import "net/http"
import "fmt"

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {
	fmt.Println("开始启动～")
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":8080", nil)
}
