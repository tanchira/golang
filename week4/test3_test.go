package test1

import "testing"

func TestHel(t *testing.T) {
	got := Hel()
	want := "Hello World"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
