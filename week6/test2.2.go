package week6

import "fmt"

func main(number int, number2 int) int {
	sum := number - number2
	return sum
}
func week6() {
	sum := main(9, 7)
	fmt.Println(sum)
}
