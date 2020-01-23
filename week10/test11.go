package main

import "fmt"

func send1(first <-chan string) {
	first <- "hello"

}
func receivel(first <-chan string) {
	fmt.Println(<-first)
}
func send2(first chan<- string) {

}
