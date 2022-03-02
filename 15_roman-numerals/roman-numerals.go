// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/roman-numerals

// For practice about 'property-based test'. Property based tests help you do this by throwing random data at your code and verifying the rules you describe always hold true.

package roman_numeral

import (
	"strings"
)

const (
	ErrInvalidArabic = ConvertErr("Failed to convert to Roman numeral because it is invalid Arabic value")
	ErrInvalidRoman  = ConvertErr("Failed to convert to Arabic numeral because it is invalid Roman value")
)

type ConvertErr string

func (e ConvertErr) Error() string { // implement 'error' interface. : https://go.dev/blog/error-handling-and-go
	return string(e)
}

func ConvertToRoman(arabic uint16) (string, error) {
	if arabic == 0 || arabic > 3999 {
		return "", ErrInvalidArabic
	}

	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String(), nil
}

func ConvertToArabic(roman string) (total uint16) {
	for _, symbols := range windowedRoman(roman).Symbols() {
		total += allRomanNumerals.ValueOf(symbols...)
	}
	return
}

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

type romanNumerals []RomanNumeral

func (r romanNumerals) ValueOf(symbols ...byte) uint16 {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}

func (r romanNumerals) Exists(symbols ...byte) bool {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return true
		}
	}
	return false
}

var allRomanNumerals = romanNumerals{
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

type windowedRoman string // take care of extracting the numerals, offering a 'Symbols' method below to retrieve them as a slice.

func (w windowedRoman) Symbols() (symbols [][]byte) {
	for i := 0; i < len(w); i++ {
		symbol := w[i]
		notAtEnd := i+1 < len(w)

		if notAtEnd && isSubtractive(symbol) && allRomanNumerals.Exists(symbol, w[i+1]) { // the symbol we are currently dealing with is a two-character subtractive symbol.
			symbols = append(symbols, []byte{symbol, w[i+1]})
			i++
		} else {
			symbols = append(symbols, []byte{symbol})
		}
	}
	return
}

func isSubtractive(symbol uint8) bool {
	return symbol == 'I' || symbol == 'X' || symbol == 'C'
}
