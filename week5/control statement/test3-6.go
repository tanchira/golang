package main

import "fmt"

func main() {
	essay := []bool{true, false, true, true, false, true}
	for _, essay := range essay {
		fmt.Printf("%t \n", essay)
	}
}
