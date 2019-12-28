package main

type student struct {
	name  string
	age   int
	email string
}

func main() {
	std := student{name: "Goku"}
	total := &std
	(*total).age=20
	total.email = "Goku@super.saiya"