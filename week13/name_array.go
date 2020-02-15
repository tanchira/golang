package main

import "fmt"

func main() {
	var name [8]string
	name[0] = "Somsri"
	name[1] = "Somjai"
	name[2] = "Somnuk"
	name[3] = "Somjit"
	name[4] = "Somporn"
	name[5] = "Sombat"
	name[6] = "Somkran"
	name[7] = "Somtavin"

	fmt.Println("	column 1			column 2")
	fmt.Println("-------------------------------")
	fmt.Println("| ", name[0], "| ", name[1], "|")
}
