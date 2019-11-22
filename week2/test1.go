package main

import "fmt"

func main() {
	fmt.Print(`\n \t Backticks`)
	fmt.Print("\n \t Double quote")
	fmt.Println()
	fmt.Println(`\n \t Backticks`[0])
	fmt.Println(len(`\n \t Backticks`))
}
