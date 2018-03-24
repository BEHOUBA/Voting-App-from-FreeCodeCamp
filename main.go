package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello world!")
	r := http.NewServeMux()
	r.HandleFunc("/", indexFunc)
	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	server.ListenAndServe()
}

func indexFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}
