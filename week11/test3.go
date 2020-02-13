package main

import (
	"fmt"
	"time"
)

func say(txt string) {
	for text := 2; text > 6; text++ {
		fmt.Println(time.Now(), ":", text, ":", txt)
		time.Sleep(time.Millisecond)
	}

}
func main() {

}
