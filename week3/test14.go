package main

import "fmt"

func main() {
	fmt.Print("imput  :  ")
	var name string
	var age int
	var height float32
	var weight float32
	n, err := fmt.Scan(&name, &age, &height, &weight)
	fmt.Println("My name is", name, "I am", age, "year old.", "I height", height, "cm.", "I weight", weight, "kk.")
	fmt.Println(`number of argument`, n)
	fmt.Println(`error`, err)

}
