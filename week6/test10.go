package main

import "fmt"

func suntract(number int) {
	number = number - 1
}

func main() {
	sum := 10
	suntract(sum)
	fmt.Println(sum)
}
