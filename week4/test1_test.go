package test1

import "testing"

func TestAverage(t *testing.T) {
	v := Average(3)
	if v != 4 {
		t.Error("Average of 3,5 is 4, not", v)

		v := Average(5)
		if v != 4 {
			t.Error("Average of 3,5 is 4, not", v)
		}

	}
}
