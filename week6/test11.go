package main

func compute(fn func(int, int) int) int {
	return fn(3, 4)
}

func main() {
	sum := func(num1, num2 int) int {
		return num1 + num2
	}
}
