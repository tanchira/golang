package main

type student struct {
	name  string
	age   int
	email string
}

func (std student) growUp(i int) {
	std.age = std.age + 1

}
