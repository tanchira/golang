package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	// read whole file
	myText, myErr := ioutil.ReadFile(".txt")

	if myErr != nil {
		panic(myErr)

	}

	fmt.Print(string(myText))

}
