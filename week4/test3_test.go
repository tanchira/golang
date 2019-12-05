package test3

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello,Wiorld"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
