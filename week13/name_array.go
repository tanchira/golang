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

	fmt.Println("  column 1	column 2")
	fmt.Println("-------------------------------")
	fmt.Println("| ", name[0], "  | ", name[4], " |")
	fmt.Println("-------------------------------")
	fmt.Println("| ", name[1], "  | ", name[5], "  |")
	fmt.Println("-------------------------------")
	fmt.Println("| ", name[2], "  | ", name[6], " |")
	fmt.Println("-------------------------------")
	fmt.Println("| ", name[3], "  | ", name[7], "|")
	fmt.Println("-------------------------------")

}
