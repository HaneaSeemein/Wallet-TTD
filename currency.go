package Wallet

type Currency struct{
	valueInINR float64
}

var Rupee Currency = Currency{valueInINR: 1.0}
var Dollar Currency = Currency{valueInINR: 82.47}

func ConvertCurrency(from, to Currency) float64 {
  return from.valueInINR/to.valueInINR
}