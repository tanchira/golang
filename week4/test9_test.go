package test1

import "testing"

func TestNumber(t *testing.T) {
	want := 2
	got := Number(1, 1)
	if want != got {
		t.Error("error got not to want")
	}
	want := 2
	got := Number(2, 4)

}
