package fs

import (
	"errors"
	"io/ioutil"
)

type FileReader interface {
	ReadFile(filename string) ([]byte, error)
}

type File struct{}

// https://golang.org/pkg/io/ioutil/#example_ReadFile
func (f *File) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func Read(filename string, reader FileReader) (string, error) {
	if filename == "" {
		return "", errors.New("filename needs to be informed")
	}
	content, err := reader.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
