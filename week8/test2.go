package main

type student struct {
	name  string
	age   int
	email string
}

func main() {
	var std [10]student
	std[0 = student{"Goku",18,"Goku@super.saiya"}]

	fmt.Println(std[0])
	fmt.Println(std[0].name)
}
