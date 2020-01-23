package main

import "fmt"

func test(txt int, txt2 string) {
	for First := 1; First < 2; First++ {
		fmt.Println(First, ".", txt2, txt)
	}

}
func main() {
	go test("10")

}
