package main

import "fmt"

type myError struct {
	error string
}

func (e myError) Error() string {
	return e.error
}
func say(word string) error {
	if word == "hi" {
		return myError{"can't say hi"}
	}
	return nil
}
func main() {
	ever := say("hello")
	fmt.Println(ever)
	ever2 := say("hi")
	fmt.Println(ever2)
}
