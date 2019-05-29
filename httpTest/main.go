package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	url := "http://10.1.160.68:8080"
	req, err := http.NewRequest("POST", url, strings.NewReader("body"))
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	fmt.Printf("req: %v", req)
}
