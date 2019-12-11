package main

import "fmt"

//func main() {
//essay := []bool{true, false, true, true, false, true}
//for _, essay := range essay {
//	fmt.Printf("%t \n", essay)
//}
//}

//func main() {
//	essay := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//	for _, essay := range essay {
//		fmt.Printf(" \n", essay)
//	}
//}

func main() {
	essay := []string{"a", "b", "c", "d", "e"}
	for _, essay := range essay {
		fmt.Println(essay)
	}
}
