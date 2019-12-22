package main

import "fmt"

func handlepanic() {
	r := recover()
	if r == "to much" {
		fmt.Println("your number out of range")
	}
}
