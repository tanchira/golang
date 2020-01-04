package main

import "fmt"

func main() {
	Hugesales := 50000 / 100 * 30
	sales := 49999 / 100 * 10
	fmt.Println("			Report  Bonus			")
	fmt.Println("***********************************")
	var no int
	fmt.Println("  No (ลำดับพนักงาน)	=  :			 ")
	var name string
	var summit int
	fmt.Scan(&no, &name, &summit)

	fmt.Println("  No (ลำดับพนักงาน)	=  :			 ")
	fmt.Println("  Name (ชื่อพนักงาน)	 =	:     	  ")
	fmt.Println("  Summit (ยอดขาย)     =  :    ", summit)
	fmt.Println("***********************************")
	if summit >= 50000 {
		fmt.Println(" Bonus (โบนัส)         =      ", Hugesales)
	} else {
		fmt.Println(" Bonus (โบนัส)         =      ", sales)
	}
}
