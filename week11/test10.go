package main

import (
	"fmt"
	"sync"
)

func generteInt(min, max int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := min; i <= max; i++ {
		ch <- 1
	}

}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go generteInt(10, 20, ch, &wg)
	go generteInt(50, 200, ch, &wg)
	go func() {
		wg.Wait()
		close(ch)
	}()
	sum := 0
	for i := range ch {
		sum += i
	}
	fmt.Println(sum)

}
