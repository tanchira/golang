package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bs, err := ioutil.ReadFile("text.txt")
	if err != nil {
		return
	}
	srt := string(bs)
	fmt.Println(srt)
}
