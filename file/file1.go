package main

import "os"

func main() {
	file, err := os.Create("name.txt")
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString("Hello \n")
	file.WriteString("i am name.txt")
}
