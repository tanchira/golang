package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("Hello Tanchira")
	text := make([]byte, 3)
	for {
		n, err := reader.Read(text)
		if err == io.EOF {
			break
		}
		fmt.Println(string(text[:n]))
	}
}
