package main

import "fmt"

func main() {
	dawn1 := []int{1, 2, 3}
	dawn2 := dawn1
	fmt.Println(dawn1, dawn2)
	dawn1[0] = 10
	fmt.Println(dawn1, dawn2)

}
