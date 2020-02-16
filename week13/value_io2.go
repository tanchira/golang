package main

import "fmt"

func main() {
	value1 := 9
	value2 := 6
	value3 := 0

	value3 = value1 + value2
	fmt.Printf("%d %d %d\n", value1, value2, value3)
	value3 = value1 - value2
	fmt.Printf("%d %d %d\n", value1, value2, value3)
	value3 = value1
	value3 = value3 - value2
	value3 = value3 + value1

}
