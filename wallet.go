package Wallet

type Wallet struct{
	balance float64
	nativeCurrency Currency
}

func newWallet(amount float64, c Currency) *Wallet{
	if amount>0.0{
		return &Wallet{
			balance: amount,
			nativeCurrency: c, 
		}
	}
	return nil
}

func (w *Wallet) Credit(amount float64) float64 {
	if amount<=0{
		return w.balance
	}
	w.balance = w.balance+amount
	return w.balance
}

func (w *Wallet) CreditIn(c Currency, amount float64) float64 {
	if amount<=0{
		return w.balance*ConvertCurrency(w.nativeCurrency,c)
	}
	w.balance = w.balance+(amount*ConvertCurrency(c,w.nativeCurrency))
	return w.balance*ConvertCurrency(w.nativeCurrency,c)
}

func (w *Wallet) Debit(amount float64) float64 {
	if amount<=0.0{
		return -1.0
	}	
	if amount>w.balance{
		return -1.0
	}
	w.balance = w.balance-amount
	return w.balance
}

func (w *Wallet) DebitIn(c Currency, amount float64) float64 {
	amt := amount*ConvertCurrency(c,w.nativeCurrency)
	if amt<=0{
		return -1
	}	
	if amt>w.balance{
		return -1
	}
	w.balance = w.balance-amt
	return w.balance*ConvertCurrency(w.nativeCurrency,c)
}

func (w *Wallet) CheckBalance() float64 {
	return w.balance
}

func (w *Wallet) CheckBalanceIn(c Currency) float64 {
	return w.balance*ConvertCurrency(w.nativeCurrency,c)
}