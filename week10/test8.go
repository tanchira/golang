package main

import (
	"fmt"
	"sync"
	"time"
)

func incrment(data *int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	start := time.Now()
	defer wg.Done()
	defer mutex.Unlock()
	mutex.Lock()
	*data++
	fmt.Println(time.Since(start), "Increment to :", *data)
}
func main() {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	data := 10
}
