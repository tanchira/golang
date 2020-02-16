package main

import "fmt"

func main() {
	var salary int
	fmt.Println("salary : ")
	fmt.Scan(&salary)

	if salary >= 10000 {
		deduct := salary / 100 * 10
		fmt.Println("หักภาษีทั้งหมด", deduct)
	} else if salary >= 5001 {
		deduct := salary / 100 * 5
		fmt.Println("ให้หักภาษีทั้งหมด", deduct2)
	}

}
