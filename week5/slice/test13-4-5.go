package main

func main() {
	num1 := []int{1, 2, 3, 4}
	num2 := []int{5, 6, 7, 8}
	num2 = append(num1, num2...)
}
