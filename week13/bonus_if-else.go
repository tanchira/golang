package main

import "fmt"

func main() {
	Hugesales := 50000 / 100 * 30
	sales := 49999 / 100 * 10
	fmt.Println("			Report  Bonus			")
	fmt.Println("*************************************************")
	var no int
	fmt.Print("No(ลำดับพนักงาน)  :  ")
	fmt.Scan(&no)
	var name string
	fmt.Print("Name (ชื่อพนักงาน) :  ")
	fmt.Scan(&name)
	var summit int
	fmt.Print("Summit(ยอดขาย)  :  ")
	fmt.Scan(&summit)
	fmt.Println("*************************************************")
	if summit >= 50000 {
		fmt.Println(" Bonus (โบนัส)         =      ", Hugesales)
	} else {
		fmt.Println(" Bonus (โบนัส)         =      ", sales)
	}

}