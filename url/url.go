package main

import (
	"fmt"
	"net/url"
	"path/filepath"
)

func main() {

	raw := "http://domain.com:8080/some/dir/some_file.tar.gz"
	urlStruct, err := url.Parse(raw)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Hostname: %s\n", urlStruct.Hostname())
	fmt.Printf("Path: %s\n", urlStruct.Path)
	fmt.Printf("Filename: %s\n", filepath.Base(urlStruct.Path))
	fmt.Printf("Dir: %s\n", filepath.Dir(urlStruct.Path))
}
