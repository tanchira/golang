package main

import "fmt"

func makeEvan() func() int {
	even := 0
	return func() int {
		even = even + 2
		return even
	}
}

func main() {
	nextEven := makeEvan()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}
