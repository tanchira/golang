package main

import "fmt"

func main() {
	var elecbill float32
	fmt.Println("Power unit :")
	fmt.Scan(&elecbill)
	var waterbill float32
	fmt.Println("Water unit :")
	fmt.Scan(&waterbill)
	var all float32
	elecbill = elecbill * 5
	waterbill = waterbill * 2
	fmt.Println("Electricity bill", elecbill)
	fmt.Println("Water bill", waterbill)

	elecbill = elecbill + ((elecbill * 7) / 100)
	waterbill = waterbill + ((waterbill * 7) / 100)
	all = elecbill + waterbill
	fmt.Println("totel ", all)

}
