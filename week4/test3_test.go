package test1

import "testing"

func TestHi(t *testing.T) {
	got := Hi()
	want := "Hello,Wiorld"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
