package Wallet

type Currency struct{
	sign rune
	valueInINR float64
}

var Rupee Currency = Currency{sign:'₹',valueInINR: 1.0}
var Dollar Currency = Currency{sign:'$', valueInINR: 82.47}
var Euro Currency = Currency{sign:'€', valueInINR: 100}

func ConvertCurrency(from, to Currency) float64 {
  return from.valueInINR/to.valueInINR
}