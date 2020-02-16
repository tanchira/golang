package main

import "fmt"

func main() {
	var elecbill int
	fmt.Println("Power unit :")
	fmt.Scan(&elecbill)
	var waterbill int
	fmt.Println("Water unit :")

	var all int
	elecbill = elecbill * 5
	waterbill = waterbill * 2
	fmt.Println("Water bill", waterbill)
	fmt.Println("Electricity bill", elecbill)
	elecbill = elecbill + ((elecbill * 7) / 100)
	waterbill = waterbill + ((waterbill * 7) / 100)
	all = elecbill + waterbill
	fmt.Println("totel ", all)

}
