package main

import "fmt"

func main() {
	Hugesales := 50000 / 100 * 30
	sales := 49999 / 100 * 10
	fmt.Println("			Report  Bonus			")
	fmt.Println("***********************************")
	var no int
	fmt.Print("  No (ลำดับพนักงาน)	=	")
	var name string
	fmt.Print("  Name (ชื่อพนักงาน)	 =	")
	var summit int
	fmt.Print("  Summit (ยอดขาย)	=	", summit)
	fmt.Scan(&no, &name, &summit)

	fmt.Println("***********************************")
	if summit >= 50000 {
		fmt.Println(" Bonus (โบนัส)         =      ", Hugesales)
	} else {
		fmt.Println(" Bonus (โบนัส)         =      ", sales)
	}
}
