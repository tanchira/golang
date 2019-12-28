package main

import "fmt"

type student struct {
	name  string
	age   int
	email string
}

func main() {
	var all student
	all.name = "Goku"
	all.age = 18
	all.email = "Goku@super.saiya"

	all2 := student{"Gohan", 3, "Gohan@super.saiya"}
	all3 := student{name: "videl", email: "Videl@daughter.satan"}
	all4 := student{age: 20}

	fmt.Println(all)
	fmt.Println(all2)
	fmt.Println(all3)
	fmt.Println(all4)

}
