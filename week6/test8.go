package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("before")
		continue
		fmt.Println("after")
	}
	fmt.Println("next statment")
}
