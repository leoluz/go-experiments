package main

import (
	"fmt"
	"html"
	"net/http"
)

const (
	Countries string = "/countries/"
)

func main() {
	http.HandleFunc("/", helloWorldHandler)
	http.HandleFunc(Countries, countriesHandler)
	http.ListenAndServe(":8080", nil)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path[1:]))
}

func countriesHandler(w http.ResponseWriter, r *http.Request) {
	country := r.URL.Path[len(Countries):]
	fmt.Fprintf(w, "Hello from %v", country)
}
