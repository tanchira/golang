package main

import "fmt"

func main() {
	fmt.Printf("My name is %s, I am %d years old \n", "bam", 19)
	fmt.Printf("type = %T \n", 11)
	fmt.Printf("type = %T \n", 3.10159265359)
	fmt.Printf("Pi =%2.2f \n", 3.10159265359)
	fmt.Printf("Pi =%9.f \n", 3.10159265359)
	fmt.Printf("Pi =%-9.f \n", 3.10159265359)
	fmt.Printf("Pi =%09.f \n", 3.10159265359)
	fmt.Printf("Pi =%09.2f \n", 3.10159265359)
	fmt.Printf("Pi =%E \n", 3.10159265359)

	fmt.Printf("1 > 2 =%t \n", 1 > 2)
	fmt.Printf("10 (base 2) =%b \n", 10)
}
