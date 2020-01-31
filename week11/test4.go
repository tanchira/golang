package main

import "io/ioutil"

func main() {
	bs, err := ioutil.ReadFile("text.txt")
	if err != nil {
		return
	}
	srt := string(bs)
}
