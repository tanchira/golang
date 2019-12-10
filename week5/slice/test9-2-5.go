package main

import "fmt"

func main() {
	name := [5]string{"tanhira"}
	name[5] = "tanchira"
	fmt.Println(name)
}
