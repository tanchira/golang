package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("HelloTanchira")
	text := make([]byte, 2)
	for {
		n, err := reader.Read(text)
		if err == io.EOF {
			break
		}
		fmt.Println(string(text[:n]))
	}
}
