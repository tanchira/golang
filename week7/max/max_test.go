package main

import (
	"testing"
)

func TestMax(t *testing.T) {
	want := 99
	max := Max()
	if want != max {
		t.Error("want false but got true")
	}

}
