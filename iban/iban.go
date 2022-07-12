package main

import (
	"errors"
	"strconv"
	"unicode"
	"bytes"
	"math/big"
)

const IbanMaxLength = 34
const MapLetter2DigitRef = 55
const (
	Modulus = 97
	CorrectRemainder = 1
	Base = 10
)

func validateIban(iban string) error {
	//order is important!!!
	validationFuncs := [](func(string) error){
		checkIbanMaxLength,
		checkIbanAtLeast4char,
		checkIbanCountryCode,
		checkIbanCorrectLength,
		checkIbanCheckDigitAreDigits,
		checkIbanHasValidChar,
		checkIbanWithCheckDigits}

	for _, fun := range validationFuncs {
		if err := fun(iban); err != nil {
			return err
		}
	}
	return nil
}

func checkIbanMaxLength(iban string) error {
	if len(iban) > IbanMaxLength {
		return errors.New("iban has to many characters")
	}
	return nil
}

func checkIbanAtLeast4char(iban string) error {
	if len(iban) < 4 {
		return errors.New("iban has to few characters")
	}
	return nil
}

// Not safe to call before checkIbanAtLeast4char
func checkIbanCountryCode(iban string) error {
	countryCode := getCountryFromIban(iban)
	if _, exists := getCountriesIbanInfo()[countryCode]; !exists {
		return errors.New("iban does not have a valid country code")
	}
	return nil
}

// Not safe to call before checkIbanCountryCode
func checkIbanCorrectLength(iban string) error {
	countryCode := getCountryFromIban(iban)
	countryIbanIfo, _ := getCountriesIbanInfo()[countryCode];
	if len(iban) != countryIbanIfo.Length {
		return errors.New("iban does not have a correct length")
	}
	return nil
}

// Not safe to call before checkIbanAtLeast4char
func checkIbanCheckDigitAreDigits(iban string) error {
	if _, err := strconv.Atoi(getCheckDigitsFromIban(iban)); err != nil {
		return errors.New("iban's check digit are not digits")
	}
	return nil
}

func checkIbanHasValidChar(iban string) error {
    for _, c := range iban {
        if !isValidIbanChar(c) {
			return errors.New("iban does contain invalid characters")
        }
    }
    return nil
}

func checkIbanWithCheckDigits(iban string) error {
	number, ok := new(big.Int).SetString(replaceLetterWithNumber(moveFirst4char2Last(iban)), Base)
	if !ok {
		return errors.New("Convert string to big.Int failed when validating iban with check digit")
	}

	remainder := new(big.Int)
	if remainder.Mod(number, big.NewInt(Modulus)).Cmp(big.NewInt(CorrectRemainder)) != 0 {
		return errors.New("iban is invalid according to check digits")
	}
	return nil
}

func moveFirst4char2Last(iban string) string {
	return iban[4:] + iban[0:4]
}

func replaceLetterWithNumber(str string) string {
	var buffer bytes.Buffer

	for _, c := range str {
		if unicode.IsUpper(c) {
			buffer.WriteString(mapLetter2Digits(c))
		} else {
			buffer.WriteString(string(c))
		}
	}
	return buffer.String()
}

// A -> 10, B -> 11, C -> 12, ..., Y -> 34, Z -> 35
func mapLetter2Digits(letter rune) string {
	return strconv.Itoa(int(letter) - MapLetter2DigitRef)
}

func isValidIbanChar(c rune) bool {
    return unicode.IsDigit(c) || unicode.IsUpper(c)
}

// Not safe to call before checkIbanAtLeast4char
func getCountryFromIban(iban string) string {
	return iban[0:2]
}

// Not safe to call before checkIbanAtLeast4char
func getCheckDigitsFromIban(iban string) string {
	return iban[2:4]
}