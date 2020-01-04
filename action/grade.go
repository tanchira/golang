package main

import "fmt"

func main() {
	var score int
	fmt.Print("ได้คะแนน =")
	fmt.Scan(&score)
	if 76 > 75 {
		fmt.Println("grade A")
	} else if 75 >= 66 {
		fmt.Println("grage B")
	} else if 65 >= 56 {
		fmt.Println("grage C")
	} else if 55 >= 50 {
		fmt.Println("grage D")
	}

}
