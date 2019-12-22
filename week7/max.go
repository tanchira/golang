package main

import "fmt"

func handlePanic() {

	Max := recover()
	if Max == "to much" {
		fmt.Println("error")
		main()
	}
}
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
	panic("not error")
	max := Max(1, 2, 3, 4, 9, 7, 8)
	fmt.Println(max)

}
