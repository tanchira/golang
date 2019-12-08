package test1

import "testing"

func TestNumber(t *testing.T) {
	want := 5
	got := (8,13)
	if want != got {
		t.Error("got 8 and 13 not 5",want,got)
	}

}
