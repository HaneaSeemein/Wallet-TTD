package Wallet

type Currency struct{
	valueInINR float
}

func ConvertCurrency(from, to Currency) float {
	return from.valueInINR/to.valueInINR
}

Rupee := Currency{valueInINR: 1}
Dollar := Currency{valueInINR: 82.47}

type Wallet struct{
	var balance float
	nativeCurrency Currency
}

func newWallet[T any](amt T, c Currency) *Wallet{
	amount:=float(amt)
	if amount<0{
		return nil
	}
	if !c{
		return nil
	}
	return &Wallet{
		balance: amount,
		nativeCurrency: c, 
	}
}

func (w Wallet) Credit[T any](amt T) float {
	amount:=float(amt)
	if amount<=0{
		return w.balance
	}
	w.balance = w.balance+amount
	return w.balance
}

func (w Wallet) CreditIn[T any](c Currency, amt T) float {
	amount:=float(amt)
	if amount<=0{
		return w.balance*ConvertCurrency(w.nativeCurrency,c)
	}
	w.balance = w.balance+(amount*ConvertCurrency(c,w.nativeCurrency))
	return w.balance*ConvertCurrency(w.nativeCurrency,c)
}

func (w Wallet) Debit[T any](amt T) float {
	amount:=float(amt)
	if amount<=0 || amount<w.balance{
		return -1
	}
	w.balance = w.balance-amount
	return w.balance
}

func (w Wallet) DebitIn[T any](c Currency, amt T) float {
	amount:=float(amt)
	amt := amount*ConvertCurrency(c,w.nativeCurrency)
	if amt<=0 || amt<w.balance{
		return -1
	}
	w.balance = w.balance-amt
	return w.balance*ConvertCurrency(w.nativeCurrency,c)
}

func (w Wallet) CheckBalance() float {
	return w.balance
}

func (w Wallet) CheckBalanceIn(c Currency) float {
	return w.balance*ConvertCurrency(w.nativeCurrency,c)
}