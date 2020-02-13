package main

import (
	"fmt"
	"sync"
	"time"
)

func increment(data *int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	start := time.Now()
	defer wg.Done()
	defer mutex.Unlock()
	mutex.Lock()
	*data++
	fmt.Println(time.Since(start), "Increment to:", *data)

}
func resd(data *int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	start := time.Now()
	defer wg.Done()
	defer mutex.Unlock()
	mutex.Lock()
	fmt.Println(time.Since(start), "Data =", *data)

}

func main() {

}
