package main

import (
	"fmt"
	"reflect"
)

func main() {
	num1 := []int{1, 2, 3}
	num2 := []int{1, 2, 3}
	fmt.Println(reflect.DeepEqual(num1, num2))
	num3 := []string{"hi", "hello"}
}
