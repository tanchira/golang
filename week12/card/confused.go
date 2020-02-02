package card

import (
	"fmt"
	"math/rand"
)

func card() {
	num := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	henddis := make([]string, 6)
	for i := 0; i < 5; i++ {
		hend := num[rand.Intn(len(num))]
		henddis = append(henddis, hend)
	}
	fmt.Println(henddis)
}
func pick() {
	num := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	newcard := fmt.Sprint(num[rand.Intn(len(num))])

}
