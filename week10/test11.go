package main

func send1(first <-chan string) {
	first <- "hello"

}
func receivel(first <-chan string) {

}
