package card

import (
	"fmt"
	"math/rand"
)

func main() int {
	var pair int
	num := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	ren := make([]string, 6)
	for i := 0; i < 5; i++ {
		hend := num[rand.Intn(len(num))]
		ren = append(ren, hend)
	}
	fmt.Println(ren)
	fmt.Println(pair)

}

func cke() {

}
