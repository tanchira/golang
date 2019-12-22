package main

import "fmt"

func main() {
	var number int
	fmt.Println("type number :")
	_, e := fmt.Scan(&number)
	if e != nil {
		panic("to much")
	}
	fmt.Println("your number :", number)
}
