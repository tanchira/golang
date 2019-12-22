package main

import "fmt"

func handlepanic() {
	text := recover()
	fmt.Println(text)
}
