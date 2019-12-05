package test2

func TestHello(t. *testing.T) {
	output := ("Hello ,World")
	got := ("Hello world")
	if got != output{
		t.Error("error")
	}
	
}
