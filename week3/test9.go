package main

import "fmt"

func main() {
	n, e := fmt.Print("tanchira", "Paewkrathok", 123, 456, 789, 123, 456, 789, "\n")
	fmt.Print(" munber of bytes written :", n, "\n")
	fmt.Print("write error encountered :", e, "\n")
}
