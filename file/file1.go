package main

import "os"

func main() {
	file, err := os.Create("myfile.txt")
	if err != nil {
		return
	}
}
