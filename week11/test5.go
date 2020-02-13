package main

import (
	"fmt"
	"sync"
	"time"
)

func say(txt string, sleep time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(txt)
	time.Sleep(time.Millisecond * sleep)

}
func main() {

}
