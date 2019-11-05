package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	url := "http://httpbin.org/status/200"
	c := &http.Client{Timeout: time.Second * 10}
	status, err := DoGet(c, url)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	fmt.Printf("Status: %d", status)
}

func DoGet(client *http.Client, url string) (int, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("error sending request: %s", err)
	}
	return resp.StatusCode, nil
}
