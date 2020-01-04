package main

import "fmt"

func main() {

	deduct := 10000 / 100 * 10
	deduct2 := 5001 / 100 * 5
	if salary >= 10000 {
		fmt.Println("salary : ")
		var salary int
		fmt.Scan(&salary)
	} else if salary >= 5001 {
		fmt.Println(deduct2, "ให้หักภาษีอัตรา 5% ของเงินเดือน")
	}
}
