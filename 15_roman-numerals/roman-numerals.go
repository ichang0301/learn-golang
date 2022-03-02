package roman_numeral

import (
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}

var allRomanNumerals = RomanNumerals{
	{Value: 1000, Symbol: "M"},
	{Value: 900, Symbol: "CM"},
	{Value: 500, Symbol: "D"},
	{Value: 400, Symbol: "CD"},
	{Value: 100, Symbol: "C"},
	{Value: 90, Symbol: "XC"},
	{Value: 50, Symbol: "L"},
	{Value: 40, Symbol: "XL"},
	{Value: 10, Symbol: "X"},
	{Value: 9, Symbol: "IX"},
	{Value: 5, Symbol: "V"},
	{Value: 4, Symbol: "IV"},
	{Value: 1, Symbol: "I"},
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) int {
	total := 0
	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		// look ahead to next symbol if we can and the current symbol is base 10 (only valid subtractors)
		if couldBeSubtractive(i, symbol, roman) {
			if value := allRomanNumerals.ValueOf(symbol, roman[i+1]); value != 0 { // get the value of the two character string
				total += value
				i++ // move past this character too for the next loop
			} else {
				total += allRomanNumerals.ValueOf(symbol)
			}
		} else {
			total += allRomanNumerals.ValueOf(symbol)
		}
	}
	return total
}

func couldBeSubtractive(index int, currentSymbol uint8, roman string) bool {
	isSubtractiveSymbol := currentSymbol == 'I' || currentSymbol == 'X' || currentSymbol == 'C'
	return index+1 < len(roman) && isSubtractiveSymbol
}
