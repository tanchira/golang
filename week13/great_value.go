package main

import "fmt"

func main() {
	var num [10]int
	var i, j, max, min []int

	for i := 0; i < 10; i++ {
		fmt.Println("Enter number :")
		fmt.Scanf("%d", &num[i])
	}

}
