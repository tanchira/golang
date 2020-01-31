package main

import (
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("Hello")
	index := strings.NewReader("Tanchira")
	text := make([]byte, 2)
	for {
		n, err := reader.Read(text)
		if err == io.EOF {
			break
		}
	}

}
