package main

import "fmt"

func main() {
	var value1, value2 int
	var value3, value4 float64

	value1 = 10
	value3 = 5.5
	value2 = 3
	value4 = 3.5
	fmt.Printf("%d \n", value1+value2)
	fmt.Printf("%d \n", value1-value2)
	fmt.Printf("%d \n", value1*value2)
	fmt.Printf("%d \n", value1/value2)
	fmt.Printf("%d \n", value1%value2)
	fmt.Printf("%d \n", value3+value4)
	fmt.Printf("%d \n", value3-value4)
	fmt.Printf("%d \n", value3*value4)
	fmt.Printf("%d \n", value3/value4)
	fmt.Printf("%d \n", value3%value4)
}
