package main

func swap(num1, num2 int) (int, int) {
	return num2, num1
}

func main() {
	num1, num2 := swap(10, 5)
}
