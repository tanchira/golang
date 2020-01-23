package main

import (
	"fmt"
	"time"
)

func main() {
	data := 20
	go func() {
		data = 10
	}()

	go func() {
		fmt.Println(data)

	}()
	time.Sleep(time.Millisecond)
}
