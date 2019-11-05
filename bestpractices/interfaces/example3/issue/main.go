package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	content, err := Read("/etc/hosts")
	if err != nil {
		fmt.Printf("error reading file: %s", err)
		os.Exit(1)
	}
	fmt.Printf("%s", content)
}

func Read(filename string) (string, error) {
	if filename == "" {
		return "", errors.New("filename needs to be informed")
	}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
