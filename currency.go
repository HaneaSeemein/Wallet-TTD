package Wallet

type Currency struct{
	valueInINR float
}

func ConvertCurrency(from, to Currency) float {
	return from.valueInINR/to.valueInINR
}

Rupee := Currency{valueInINR: 1}
Dollar := Currency{valueInINR: 82.47}