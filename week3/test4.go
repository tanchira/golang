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
	fmt.Printf("%04t \n", 15 > 14)
	fmt.Printf("%08t \n	", 16 > 15)
	fmt.Printf("%-9d \n", 10)
	fmt.Printf("%-1d \n", 17)
	fmt.Printf("%-2d \n", 17)
	fmt.Printf("%-3d \n", 17)
	fmt.Printf("%-4d \n", 17)
	fmt.Printf("%-8d \n", 17)
	fmt.Printf("%-9d \n", 17)
	fmt.Printf("%-9d \n", 10)
	fmt.Printf("%-9d", 10)

}
