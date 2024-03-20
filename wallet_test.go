package Wallet

import 	"testing"

func TestNewWallet(t *testing.T){
	mywallet := newWallet(Rupee, 10.0)
	got := mywallet.balance
	want := float64(10)

	if got!=want{
		t.Errorf("wallet not created")
	}
}

func TestNewWalletWithNegativeBalance(t *testing.T){
	mywallet := newWallet(Rupee, -2.0)
	got:=mywallet

	if got!=nil{
		t.Errorf("Wallet with negative balance created")
	}
}

func TestCreditAmount(t *testing.T){
	mywallet := newWallet(Rupee, 10.0)
	got := mywallet.Credit(2.0)
	want := 12.0

	if got!=want{
		t.Errorf("Amount not credited, got: %f",got)
	}
}

func TestCreditAmountIn(t *testing.T){
	mywallet := newWallet(Rupee, 50.0)
	_= mywallet.CreditIn(Dollar, 1.0)
	got:=mywallet.CheckBalance()
	want := 132.47

	if got!=want{
		t.Errorf("Amount not credited, got: %f",got)
	}
}

func TestMultipleCreditAmountIn(t *testing.T){
	mywallet := newWallet(Rupee, 82.47)
	_= mywallet.CreditIn(Dollar, 1.0)
	_= mywallet.CreditIn(Rupee, 164.94)
	got:=mywallet.CheckBalanceIn(Dollar)
	want := 4.0

	if got!=want{
		t.Errorf("Amount not credited, got: %f",got)
	}
}


func TestCreditNegativeAmount(t *testing.T){
	mywallet := newWallet(Rupee,10.0)
	got := mywallet.Credit(-4.0)
	want := 10.0

	if got!=want{
		t.Errorf("Negative amount credited, got: %f",got)
	}	
}

func TestDebitAmount(t *testing.T){
	mywallet := newWallet(Rupee, 10.0)
	got := mywallet.Debit(2.0)
	want := 8.0

	if got!=want{
		t.Errorf("Amount not debited, got: %f",got)
	}
}

func TestDebitNegativeAmount(t *testing.T){
	mywallet := newWallet(Rupee, 10.0)
	got := mywallet.Debit(-4.0)
	want := -1.0

	if got!=want{
		t.Errorf("Negative amount debited, got: %f",got)
	}	
}

func TestDebitLimitExceed(t *testing.T){
	mywallet := newWallet(Rupee, 10.0)
	got := mywallet.Debit(12.0)
	want := -1.0

	if got!=want{
		t.Errorf("Wallet has negative balance, got: %f",got)
	}	
}

func TestCheckBalance(t *testing.T){
	mywallet := newWallet(Rupee, 10.0)
	want:= mywallet.Debit(3.0)

	got := mywallet.CheckBalance()

	if got!=want{
		t.Errorf("Wrong balance, got: %f",got)
	}
}

func TestCheckBalanceIn(t *testing.T){
	mywallet := newWallet(Dollar, 10.0)
	_= mywallet.Debit(3.0)

	got := mywallet.CheckBalanceIn(Rupee)
	want := 577.29

	if got!=want{
		t.Errorf("Wrong balance in rupees, got: %f",got)
	}
}