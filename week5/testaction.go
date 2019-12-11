package main

import "fmt"

func main() {
	x := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(x[2])
	fmt.Println(x[1:3])
	fmt.Println(len(x))
}
