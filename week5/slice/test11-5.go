package main

import "fmt"

func main() {
	alphabets := [4]string{"a", "b", "c", "d"}
	fmt.Println(alphabets)

	x := alphabets[0:2]
	fmt.Println(x)

	y := x[2:4]

}
