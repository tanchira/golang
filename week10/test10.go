package main

func sum(c chan int, number ...int) {
	sum := 0
	for _, v := range number {
		sum = sum + v
	}
	c <- sum

}
func printer(c chan int) {
	go printer(c)
	go printer(c)
	go sum(c, 1, 2, 3)
	go sum(c, 10, 11)

}
