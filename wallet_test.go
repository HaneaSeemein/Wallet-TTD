package Wallet

import "testing"

func TestNewWallet(t *testing.T){
	mywallet := newWallet(10.0, Rupee)
	got := mywallet.balance
	want := float64(10)

	if got!=want{
		t.Errorf("wallet not created")
	}
}

func TestNewWalletWithNegativeBalance(t *testing.T){
	mywallet := newWallet(-2.0,Rupee)
	got:=mywallet

	if got!=nil{
		t.Errorf("Wallet with negative balance created")
	}
}

func TestNewWalletWithInvalidCurrency(t *testing.T){
	mywallet := newWallet(1232.0,babushka)
	got:=mywallet

	if got!=nil{
		t.Errorf("Unrecognized currency")
	}
}

func TestCredit(t *testing.T){
	mywallet := newWallet(10.0, Rupee)
	got := mywallet.Credit(2.0)
	want := 12.0

	if got!=want{
		t.Errorf("Amount not credited")
	}
}

func TestCreditNegativeAmount(t *testing.T){
	mywallet := newWallet(10.0, Rupee)
	got := mywallet.Credit(-4.0)
	want := 10.0

	if got!=want{
		t.Errorf("Negative amount credited")
	}	
}

func TestDebit(t *testing.T){
	mywallet := newWallet(10.0, Rupee)
	got := mywallet.Debit(2.0)
	want := 8.0

	if got!=want{
		t.Errorf("Amount not debited")
	}
}

func TestDebitNegativeAmount(t *testing.T){
	mywallet := newWallet(10.0, Rupee)
	got := mywallet.Debit(-4.0)
	want := -1.0

	if got!=want{
		t.Errorf("Negative amount debited")
	}	
}

func TestDebitLimitExceed(t *testing.T){
	mywallet := newWallet(10.0, Rupee)
	got := mywallet.Debit(12.0)
	want := -1.0

	if got!=want{
		t.Errorf("Wallet has negative balance")
	}	
}

func TestCheckBalance(t *testing.T){
	mywallet := newWallet(10.0, Rupee)
	_= mywallet.Debit(3.0)

	got := mywallet.CheckBalance()
	want := 7.0

	if got!=want{
		t.Errorf("Wrong balance")
	}
}

func TestCheckBalanceIn(t *testing.T){
	mywallet := newWallet(10.0, 	Dollar)
	_= mywallet.Debit(3.0)

	got := mywallet.CheckBalanceIn(Rupee)
	want := 577.29

	if got!=want{
		t.Errorf("Wrong balance in rupees")
	}
}