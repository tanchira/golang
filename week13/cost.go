package main

import "fmt"

func main() {
	fmt.Println("Water unit :")
	fmt.Println("Power unit :")
	var elecbill int
	var waterbill int
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
