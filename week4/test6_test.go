package test1

import "testing"

func TestPlus(t *testing.T) {
	want := Plus(2 + 3)
	got := 5

	if want != got {
		t.Error("want not equal to got ")
	}
}
