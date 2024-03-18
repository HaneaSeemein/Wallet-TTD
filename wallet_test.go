package Wallet

import "testing"

func TestRupeeToDollar(t *testing.T){
	got := RupeeToDollar(82.47)
	want := 1

	if got!=want{
		t.Errorf("Wrong conversion")
	}
}