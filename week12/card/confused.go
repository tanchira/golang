package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var pair int
	num := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	henddis := make([]string, 6)
	for i := 0; i < 5; i++ {
		hend := num[rand.Intn(len(num))]
		henddis = append(henddis, hend)
	}
	fmt.Println(henddis)
	pick()
	fmt.Print(pair)
	if pair == 3 {

	}
}
func pick() {
	num := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	newcard := fmt.Sprint(num[rand.Intn(len(num))])
	fmt.Println(newcard)

}

func card() {
	henddis := make([]string, 6)
	var pair int
	if pair == 3 {
		var throw int
	}
}
