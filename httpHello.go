package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello %s!</h1>", r.RemoteAddr)
}

func user_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello %s!</h1>", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/user", user_handler)
	http.ListenAndServe(":8080", nil)
}
