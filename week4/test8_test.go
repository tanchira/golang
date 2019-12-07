package test1

import "testing"

func TestSum(t *testing.T) {
	t.Run("Hello World", func(t *tesing.T) {
		want := true
		got := Sum("Hello World")
		if want != got {
			t.Error("want true but got fales", want, got)
		}

	})
	t.Run("Heppy birthday", func(t *testing.T) {

	})

}
