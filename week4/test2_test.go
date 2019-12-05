package test1

import "testing"

func TestHello(t *testing.T) {
	output := ("Hello ,World")
	got := ("Hello ,World")
	if got != output {
		t.Error("error")
	}

}
