package test1

func TestNumber(t *tesing.T) {
	want := 2
	got := Number(1, 1)
	if want != got {
		t.Error("error got not to want")
	}

}
