package main

import "fmt"

func subtract(x int, y int) int {
	ans := y - x
	return ans
}
func main() {
	x := subtract(10, 5)
	fmt.Println(x)
}
