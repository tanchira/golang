package main

import "os"

func main() {
	file, err := os.Create("My.txt")
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString("Hello \n")
	file.WriteString("my name is :  tanchira peawkrathok  \n")

}
