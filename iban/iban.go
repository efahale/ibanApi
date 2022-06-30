package main

import (
	"errors"
	"strconv"
)

const IbanMaxLength = 34

type countryIbanInfo struct {
    countryCode string
    length  int
}

func getCountriesIbanInfo() map[string]countryIbanInfo {
	countriesIbanInfo := []countryIbanInfo{{"AL", 28}, {"AD", 24}}
	m := make(map[string]countryIbanInfo)

	for _, countryIbanInfo := range countriesIbanInfo {
        m[countryIbanInfo.countryCode] = countryIbanInfo
    }
	return m
}

func validateIban(iban string) error {
	//order is important!!!
	validationFuncs := [](func(string) error){
		checkIbanMaxLength,
		checkIbanAtLeast4char,
		checkIbanCountryCode,
		checkIbanCorrectLength,
		checkIbanCheckDigitAreDigits,
		checkIbanHasValidChar}

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
	if len(iban) != countryIbanIfo.length {
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

func isValidIbanChar(c rune) bool {
    return ('0' <= c && c <= '9') || ('A' <= c && c <= 'Z')
}

// Not safe to call before checkIbanAtLeast4char
func getCountryFromIban(iban string) string {
	return iban[0:2]
}

// Not safe to call before checkIbanAtLeast4char
func getCheckDigitsFromIban(iban string) string {
	return iban[2:4]
}