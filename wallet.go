package Wallet

type Currency struct{
	valueInINR float
}

// rupeecoin := Currency{valueInINR: 1}
// dollarcoin := Currency{valueInINR: 82.47}

func createCurrency(value float) Currrency{
	return Currency{valueInINR: value}
}

func initialiseCurrencies(){
	const rupeecoin := createCurrency(1)
	const dollarcoin := createCurrency(82.47)
}

type Wallet struct{
	var balanceInINR float
}

func newWallet(amount float){
	return Wallet{balanceInINR: amount}
}

func ConvertCurrency(from, to Currency) float {
	return from.valueInINR/to.valueInINR
}

func Credit(c Currency, amount float){

}

func RupeeToDollar(amount float){
	return amount*ConvertCurrency(rupeecoin, dollarcoin)
}

func DollarToRupee(amount float){
	return amount*ConvertCurrency(rupeecoin, dollarcoin)
}