package main

import "fmt"

func main() {
	var num [10]int
	var i, j, max, min []int

	for i := 0; i < 10; i++ {
		fmt.Println("Enter number :")
		fmt.Scanf("%d", &num[i])
	}
min = num[0]

	for j := 0; j <= 10; j++ {
		if min > num[j] {
			min = num[j]
		}else  
		max < num[j] {
			max = num[j]
		}else
		
	}

}
