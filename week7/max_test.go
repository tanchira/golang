package main

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	max := Max(1, 2, 3, 4, 9, 7, 8)
	fmt.Println(max)
}
