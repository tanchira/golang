package main

import "fmt"

func main() {
	var score [5]int
	score[0] = 90
	score[1] = 67
	score[2] = 82
	score[3] = 71
	score[4] = 89

	for loopindex := 0; loopindex < 5; loopindex++ {
		fmt.Println("studen's score[%d]\n", score[loopindex])
	}
	return
}
