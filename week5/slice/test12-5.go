package main

import "fmt"

func main() {
	alphabets := [4]string{"a", "b", "c", "d"}
	x := alphabets[:]
	fmt.Println(x)
}
