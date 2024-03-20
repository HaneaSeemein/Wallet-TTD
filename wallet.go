package Wallet

import "fmt"

type Wallet struct{
	balance float64
	nativeCurrency Currency
	history []string
}

func (w *Wallet) EnterLog(operation string, deets string){
	
}

func newWallet(c Currency, amount float64) *Wallet{
	initlog := fmt.Sprintf("Wallet created with %s %f",string(c.sign),amount)
	if amount>=0.0{
		return &Wallet{
			balance: amount,
			nativeCurrency: c, 
			history: []string{initlog},
		}
	}
	return nil
}

func (w *Wallet) Credit(amount float64) float64 {
	if amount<=0{
		return w.balance
	}
	w.balance = w.balance+amount
	log := fmt.Sprintf("Credited: %s %f",string(w.nativeCurrency.sign),amount)
	_=append(w.history,log)
	return w.balance
}

func (w *Wallet) CreditIn(c Currency, amount float64) float64 {
	if amount<=0{
		return w.balance*ConvertCurrency(w.nativeCurrency,c)
	}
	w.balance = w.balance+(amount*ConvertCurrency(c,w.nativeCurrency))	
	log := fmt.Sprintf("Credited: %s %f",string(c.sign),amount)
	_=append(w.history,log)
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
	log := fmt.Sprintf("Debited: %s %f",string(w.nativeCurrency.sign),amount)
	_=append(w.history,log)	
	return w.balance
}

func (w *Wallet) DebitIn(c Currency, amount float64) float64 {
	amt := amount*ConvertCurrency(c,w.nativeCurrency)
	if amt<=0{
		return -1.0
	}	
	if amt>w.balance{
		return -1.0
	}
	w.balance = w.balance-amt
	log := fmt.Sprintf("Debited: %s %f",string(c.sign),amount)
	_=append(w.history,log)	
	return w.balance*ConvertCurrency(w.nativeCurrency,c)
}

func (w *Wallet) CheckBalance() float64 {
	_=append(w.history,"Balance checked")	
	return w.balance
}

func (w *Wallet) CheckBalanceIn(c Currency) float64 {
	_=append(w.history,"Balance checked")	
	return w.balance*ConvertCurrency(w.nativeCurrency,c)
}

func (w *Wallet) ShowHistory() []string{
	return w.history
}