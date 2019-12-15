package main

func swap(sum1, sum2, sum3 int) (int, int, int) {
	return sum3, sum2, sum1
}

func main() {
	sum1, sum2, sum3 := swap(30, 20, 10)
}
