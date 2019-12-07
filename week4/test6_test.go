package test1

import "testing"

func TestFind(t *testing.T) {
	t.Run("Hello World", func(t *testing.T) {
		want := true
		got := Find("Hello World")
		if got != want {
			t.Errorf("\n got %v \n want %v", want, got)
		}
	})

}
