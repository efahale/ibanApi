package main

import "testing"


func TestGetCountriesIbanInfo(t *testing.T) {

    got := getCountriesIbanInfo()
	validCountryCode := "AL"
	invalidCountryCode := "ZZ"
    if _, exists := got[validCountryCode]; !exists {
        t.Errorf("got no entry for %q, wanted one entry", validCountryCode)
    }

	if _, exists := got[invalidCountryCode]; exists {
        t.Errorf("got 1 entry for %q, wanted no entry", invalidCountryCode)
    }
}

func TestCheckIbanCountryCode(t *testing.T) {
	validIbansForThisTest := []string{
		"AL",
		"AD1400080001001234567890"} 
	for _, iban := range validIbansForThisTest {
		if got := checkIbanCountryCode(iban); got != nil {
			t.Errorf("valid iban for this test %q didn't pass validation", iban)	
		}
	}

	inValidIbansForThisTest := []string{"ZZ", "al", "aD1400080001001234567890"} 
	for _, iban := range inValidIbansForThisTest {
		if got := checkIbanCountryCode(iban); got == nil {
			t.Errorf("invalid iban for this test %q did pass validation", iban)	
		}
	}	
}

func TestCheckIbanCheckDigitAreDigits(t *testing.T) {
	validIbansForThisTest := []string{
		"AL12345678901234567890123456",
		"AB98EFGHIJKLMNOPQRSTUVWXYZ1234567890",
		"XX14",
		"AD12X"} 
	for _, iban := range validIbansForThisTest {
		if got := checkIbanCheckDigitAreDigits(iban); got != nil {
			t.Errorf("valid iban for this test %q didn't pass validation", iban)	
		}
	}

	inValidIbansForThisTest := []string{"XXXX", "AL1Q", "ADA2", "AD1&1234"} 
	for _, iban := range inValidIbansForThisTest {
		if got := checkIbanCheckDigitAreDigits(iban); got == nil {
			t.Errorf("invalid iban for this test %q did pass validation", iban)	
		}
	}	
}

func TestCheckIbanHasValidChar(t *testing.T) {
	validIbansForThisTest := []string{
		"AL12345678901234567890123456",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890",
		"AD1400080001001234567890"} 
	for _, iban := range validIbansForThisTest {
		if got := checkIbanHasValidChar(iban); got != nil {
			t.Errorf("valid iban for this test %q didn't pass validation", iban)	
		}
	}

	inValidIbansForThisTest := []string{"AL12cd", "ZZ123{", "a"} 
	for _, iban := range inValidIbansForThisTest {
		if got := checkIbanHasValidChar(iban); got == nil {
			t.Errorf("invalid iban for this test %q did pass validation", iban)	
		}
	}	
}

func TestMoveFirst4char2Last(t *testing.T) {
	want := "56781234"
	got := moveFirst4char2Last("12345678")
	if want != got {
		t.Errorf("Wanted %q, got %q", want, got)
	}
}

func TestReplaceLetterWithNumber(t *testing.T) {
	want := "101110111256"
	got := replaceLetterWithNumber("1011ABC56")
	if want != got {
		t.Errorf("Wanted %q, got %q", want, got)
	}
}

func TestValidateIban(t *testing.T) {
	validIbans := []string{"AD1400080001001234567890", "AL47212110090000000235698741"} 
	for _, iban := range validIbans {
		if got := validateIban(iban); got != nil {
			t.Errorf("valid iban %q didn't pass validation (%q)", iban, got)
		}
	}

	inValidIbans := []string{"AL", "", "AL12", 
	"AL4V212110090000000235698741",
	"ALQ7212110090000000235698741",
	"AL4721211009000000p235698741",
	"AD1400080001001234567809"}

	for _, iban := range inValidIbans {
		if got := validateIban(iban); got == nil {
			t.Errorf("invalid iban %q did pass validation", iban)	
		}
	}
}