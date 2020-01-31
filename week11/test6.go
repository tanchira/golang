package main

import "os"

func main() {
	dir, err := os.Open(".")
	if err != nil {
		return
	}
}
