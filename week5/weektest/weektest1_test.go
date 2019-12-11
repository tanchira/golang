package week5_6

import (
	"testing"
)

func TestFizzWord(t *testing.T) {
	testCases := []struct {
		input [2]int
		want  string
	}{
		{[2]int{1, 1}, "Fizz"},
		{[2]int{2, 1}, "Fizz"},
		{[2]int{2, 3}, "2"},
		{[2]int{9, 7}, "9"},
	}
	for _, test := range testCases {
		got := FizzWord(test.input[0], test.input[1])
		if got != test.want {
			t.Errorf("\n unexpected \n\t got: %v \n\t want: %v ", got, test.want)
		}
	}
}
