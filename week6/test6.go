package main

import "fmt"

func writeline(text ...interface{}) {
	for _, v := range text {
		fmt.Println(v)
	}
}

func main() {

}
