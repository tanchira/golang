package main

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

}
