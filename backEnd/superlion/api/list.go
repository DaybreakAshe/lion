package api

import (
	"fmt"
	"net/http"
)

func handleHello() {

}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")

	http.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(w, "Hello Docker")
	})

	http.HandleFunc("/list", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(w, "Hello list")
	})
}
