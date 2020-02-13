package main

import "fmt"

func Max(number ...int) int {
	max := number[0]
	for text := 0; text < len(number); text++ {
		if number[text] > max {
			max = number[text]
		}
	}
	return max
}
func main() {
	max := Max(89, 66, 11, 44, 77, 66, 24, 91, 99, 87, 39, 95, 45)
	fmt.Println(max)
	panic("Hello panic")
}
