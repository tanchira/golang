package main

import "fmt"

func main() {
	number := 8
	switch number {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	default:
		fmt.Println("unknow")
	}

}
