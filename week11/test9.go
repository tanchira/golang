package main

import (
	"fmt"
	"time"
)

func prin(tick, boom <-chan time.Time) {
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("Boom!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(20 * time.Millisecond)
		}
	}
}

func main() {
	tick := time.Tick(100 * time.Millisecond)

}
