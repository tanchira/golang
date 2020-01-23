package main

import "fmt"

func test(txt string) {
	for First := 1; First < 4; First++ {
		fmt.Println(First, txt)
	}

}
func main() {
	go test("bam")
	var input string

	fmt.Scanln(&input)
}
