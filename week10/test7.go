package main

import (
	"fmt"
	"sync"
	"time"
)

func test(txt string, sleep time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(txt)

}
