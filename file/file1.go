package main

import "os"

func main() {
	file, err := os.Create("myfile.txt")
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString("Hello \n")
}
