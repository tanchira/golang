package main

import "fmt"

type student struct {
	name  string
	age   int
	email string
}

func (std *student) growUp(i int) {
	std.age = std.age + 1

}
func main() {
	var add student
	add.age = 18
	fmt.Println(add.age)
	add.growUp(20)
	fmt.Println(add.age)
}
