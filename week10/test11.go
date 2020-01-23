package main

import "fmt"

func send1(first <-chan string) {
	first <- "Hello"

}
func receivel(first <-chan string) {
	fmt.Println(<-first)
}
func send2(first chan<- string) {
	first <- "Hello"
}
func receive2(first chan<- string) {

}
