package main

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	max := Max(89, 66, 11, 44, 77, 66, 24, 91, 99, 87, 39, 95, 45)
	fmt.Println(max)
}
