package main

import (
	"fmt"
	"sync"
	"time"
)

func test(txt string, sleep time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(txt)
	time.Sleep(time.Millisecond * sleep)

}
func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go test("bam", 1, &wg)
	go test("thanjira", 2, &wg)
	go test("Peawkrathok", 3, &wg)
	wg.Wait()
}
