package main

import "fmt"

func main() {

	salary := (int,"input :")
	deduct := 10000 / 100 * 10
	deduct2 := 5001 / 100 * 5
	if salary >= 10000 {
		fmt.Println(deduct, "ให้หักภาษีอัตรา 10% ของเงินเดือน")
	} else if salary >= 5001 {
		fmt.Println(deduct2, "ให้หักภาษีอัตรา 5% ของเงินเดือน")
	}
}
