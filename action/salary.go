package main

import "fmt"

func main() {
	var salary int
	fmt.Println("salary : ")
	fmt.Scan(&salary)
	deduct := 10000 / 100 * 10
	deduct2 := 5001 / 100 * 5
	if salary >= 10000 {
		fmt.Println("หักภาษีทั้งหมด", deduct)
	} else if salary >= 5001 {
		fmt.Println(deduct2, "ให้หักภาษีทั้งหมด", deduct2)
	}
}
