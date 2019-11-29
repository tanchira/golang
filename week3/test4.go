package main

import "fmt"

func main() {
	fmt.Printf("%+d \n", 10)
	fmt.Printf("%-d \n", 100)
	fmt.Printf("%-d \n", 10)
	fmt.Printf("%9d \n", 10)
	fmt.Printf("%15d \n", 10)
	fmt.Printf("%14d \n", 10)
	fmt.Printf("%05d \n", 10)
	fmt.Printf("%05d \n", 14)
	fmt.Printf("%010d \n", 22)
	fmt.Printf("%05g \n", 22.21)
	fmt.Printf("%04t", 15 > 14)
	fmt.Printf("%08t", 16 < 15)
}
