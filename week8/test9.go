package main

type I interface{}

func main() {
	var any I
	any = "Hello"
}
