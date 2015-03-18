package main

import (
	"bytes"
	"fmt"
	"html"
	"net/http"
	"strings"
)

const (
	Countries string = "/countries/"
)

func main() {
	http.HandleFunc("/", helloWorldHandler)
	http.HandleFunc(Countries, countriesHandler)
	http.HandleFunc(Countries+"id/", countryHandler)
	http.ListenAndServe(":8080", nil)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path[1:]))
}

func countriesHandler(w http.ResponseWriter, r *http.Request) {
	country := r.URL.Path[len(Countries):]
	fmt.Fprintf(w, "Hello from %v", country)
}

func countryHandler(w http.ResponseWriter, r *http.Request) {
	urlPaths := strings.Split(r.URL.Path, "/")
	var buffer bytes.Buffer
	for _, urlPath := range urlPaths {
		buffer.WriteString(urlPath)
		buffer.WriteString("\n")
	}
	fmt.Fprintf(w, buffer.String())
}
