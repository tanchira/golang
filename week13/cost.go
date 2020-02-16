package main

import "fmt"

func main() {
	var elecbill int
	var waterbill int
	elecbill = elecbill * 5
	waterbill = waterbill * 2
	fmt.Println("Water bill", waterbill)
	fmt.Println("Electricity bill", elecbill)
	elecbill = elecbill + ((elecbill * 7) / 100)
	waterbill = waterbill + ((waterbill * 7) / 100)
	all = elecbill + waterbill

}
