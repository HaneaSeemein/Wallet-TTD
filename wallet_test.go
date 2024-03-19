package Wallet

import "testing"

func TestNewWallet(t *testing.T){
	mywallet := newWallet(10, Rupee)
	got := mywallet.balanceInINR
	want := float(10)

	if got!=want{
		t.Errorf("wallet not created")
	}
}

func TestNewWalletWithNegativeBalance(t *testing.T){
	mywallet := newWallet(-2,Rupee)
	got:=mywallet
	want:=nil

	if got!=want{
		t.Errorf("Wallet with negative balance created")
	}
}

func TestNewWalletWithInvalidCurrency(t *testing.T){
	mywallet := newWallet(1232,babushka)
	got:=mywallet
	want:=nil

	if got!=want{
		t.Errorf("Unrecognized currency")
	}
}

func TestCredit(t *testing.T){
	mywallet := newWallet(10, Rupee)
	got := mywallet.Credit(2)
	want := 12

	if got!=want{
		t.Errorf("Amount not credited")
	}
}

func TestCreditNegativeAmount(t *testing.T){
	mywallet := newWallet(10, Rupee)
	got := mywallet.Credit(-4)
	want := 10

	if got!=want{
		t.Errorf("Negative amount credited")
	}	
}

func TestDebit(t *testing.T){
	mywallet := newWallet(10, Rupee)
	got := mywallet.Debit(2)
	want := 8

	if got!=want{
		t.Errorf("Amount not debited")
	}
}

func TestDebitNegativeAmount(t *testing.T){
	mywallet := newWallet(10, Rupee)
	got := mywallet.Debit(-4)
	want := -1

	if got!=want{
		t.Errorf("Negative amount debited")
	}	
}

func TestDebitLimitExceed(t *testing.T){
	mywallet := newWallet(10, Rupee)
	got := mywallet.Debit(12)
	want := -1

	if got!=want{
		t.Errorf("Wallet has negative balance")
	}	
}

func TestCheckBalance(t *testing.T){
	mywallet := newWallet(10, Rupee)
	balance := mywallet.Debit(3)

	got := mywallet.CheckBalance()
	want := 7

	if got!=want{
		t.Errorf("Wrong balance")
	}
}

func TestCheckBalanceIn(t *testing.T){
	mywallet := newWallet(10, 	Dollar)
	balance := mywallet.Debit(3)

	got := mywallet.CheckBalance(Rupee)
	want := 577.29

	if got!=want{
		t.Errorf("Wrong balance in rupees")
	}
}