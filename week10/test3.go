package main

import "fmt"

func test(txt int, txt2 string) {
	for First := 1; First < 2; First++ {
		fmt.Println(txt2)
	}

}
func main() {
	go test(12, "bam")
	var inputtoo string
	fmt.Scanln(&inputtoo)
}
