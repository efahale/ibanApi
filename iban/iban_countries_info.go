package main

import (
    "encoding/json"
	"log"
	"io/ioutil"
)

var ibanCountriesInfo map[string]CountryIbanInfo = map[string]CountryIbanInfo{}

type CountryIbanInfo struct {
    CountryCode string `json:"countryCode"`
    Length      int    `json:"length"`
}

func init() {
	for _, countryIbanInfo := range readCountriesIbanInfo() {
		ibanCountriesInfo[countryIbanInfo.CountryCode] = countryIbanInfo
	}
}

func getCountriesIbanInfo() map[string]CountryIbanInfo {
    return ibanCountriesInfo
}

func readCountriesIbanInfo() []CountryIbanInfo {
	content, err := ioutil.ReadFile("iban_countries_info.json")
	if err != nil {
		log.Fatal(err)
	}
	countriesIbanInfo := []CountryIbanInfo{}
	json.Unmarshal([]byte(content), &countriesIbanInfo)
	return countriesIbanInfo
}