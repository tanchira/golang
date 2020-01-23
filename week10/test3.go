package main

import "fmt"

func test(txt int, txt2 string) {
	for First := 1; First < 4; First++ {
		fmt.Println(First, txt2)
	}

}
func main() {
	go test(12, "bam")
	var inputone int
	var inputtwo string

	fmt.Scanln(&inputone, &inputtwo)
}
