package main

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	max := Max(4, 8, 9, 3, 21, 5, 10, 56, 84, 71, 32)
	fmt.Println(max)
}
