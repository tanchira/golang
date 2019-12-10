package main

import "fmt"

func main() {
	alphabets := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(alphabets)

	first := alphabets[0:3]
	fmt.Println(first)

	first2 := first[3:5]
}
