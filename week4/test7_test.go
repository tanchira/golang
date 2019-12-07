package test1

import "testing"

func TestText(t *testing.T) {
	t.Run("Happ birthday", func(t *testing.T) {
		want := false
		got := Text("Heppy birthday")
		if want != got {
			t.Errorf("want fales but got ture", want, got)
		}

	})

}
