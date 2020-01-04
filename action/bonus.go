package main

import "fmt"

func main() {
	want := 50000
	Hugesales := 50000 / 100 * 30
	sales := 49999 / 100 * 10
	fmt.Println("			Report  Bonus			")
	fmt.Println("***********************************")
	fmt.Println("  No (ลำดับพนักงาน)	=	1		 ")

	if want >= 50000 {
		fmt.Println(Hugesales)
	} else {
		fmt.Println(sales)
	}
}
