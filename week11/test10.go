package main

import "sync"

func generteInt(min, max int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := min; i <= max; i++ {

	}

}
