package main

import "os"

func main() {
	file, err := os.Create("myFile.txt")
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString("hello \n")
}
