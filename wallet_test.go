package Wallet


import "testing"

func TestRupeeToDollar(t *testing.T){
	rupee := 1

	want := 82.47
	got := RupeeToDollar(rupee)

	if(want!=got){
		t.Errorf("wrong conversion")
	}
}