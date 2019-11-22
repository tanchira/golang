package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder
	b.WriteString("Hello")
	b.WriteString("")
	b.WriteString("world")
	fmt.Println(b.String())

}
