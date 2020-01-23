package main

import "fmt"

func test(txt int) {
	for First := 1; First < 2; First++ {
		fmt.Println(First, ".", txt)
	}

}
func main() {
	go test("Hello")

}
