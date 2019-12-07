package test1

import "testing"

func TestFind(t *testing.T) {
	Run."Happ birthday",func (t *testing.T)  {
		want := false
		got :=Find("Heppy birthday")
		if want != got{
			t.Errorf("want fales but got ture")
		}
		
	}
	
}
