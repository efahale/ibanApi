package main

import (
    "sync"
)

var lock = &sync.Mutex{}
var ibanCountriesInfo *map[string]countryIbanInfo

type countryIbanInfo struct {
    countryCode string
    length  int
}

func getCountriesIbanInfo() map[string]countryIbanInfo {
    if ibanCountriesInfo == nil {
        lock.Lock()
        defer lock.Unlock()
        if ibanCountriesInfo == nil {			
			m := map[string]countryIbanInfo{}
			for _, countryIbanInfo := range countriesIbanInfo {
				m[countryIbanInfo.countryCode] = countryIbanInfo
			}
			ibanCountriesInfo = &m 	
        }
    }
    return *ibanCountriesInfo
}

var countriesIbanInfo = []countryIbanInfo{
	{"AL", 28},
	{"AD", 24},
	{"AT", 20},
	{"AZ", 28},
	{"BH", 22},
	{"BY", 28},
	{"BE", 16},
	{"BA", 20},
	{"BR", 29},
	{"BG", 22},
	{"CR", 22},
	{"HR", 21},
	{"CY", 28},
	{"CZ", 24},
	{"DK", 18},
	{"DO", 28},
	{"EG", 29},
	{"SV", 28},
	{"EE", 20},
	{"FO", 18},
	{"FI", 18},
	{"FR", 27},
	{"GE", 22},
	{"DE", 22},
	{"GI", 23},
	{"GR", 27},
	{"GL", 18},
	{"GT", 28},
	{"VA", 22},
	{"HU", 28},
	{"IS", 26},
	{"IQ", 23},
	{"IE", 22},
	{"IL", 23},
	{"IT", 27},
	{"JO", 30},
	{"KZ", 20},
	{"XK", 20},
	{"KW", 30},
	{"LV", 21},
	{"LB", 28},
	{"LY", 25},
	{"LI", 21},
	{"LT", 20},
	{"LU", 20},
	{"MT", 31},
	{"MR", 27},
	{"MU", 30},
	{"MD", 24},
	{"MC", 27},
	{"ME", 22},
	{"NL", 18},
	{"MK", 19},
	{"NO", 15},
	{"PK", 24},
	{"PS", 29},
	{"PL", 28},
	{"PT", 25},
	{"QA", 29},
	{"RO", 24},
	{"LC", 32},
	{"SM", 27},
	{"ST", 25},
	{"SA", 24},
	{"RS", 22},
	{"SC", 31},
	{"SK", 24},
	{"SI", 19},
	{"ES", 24},
	{"SD", 18},
	{"SE", 24},
	{"CH", 21},
	{"TL", 23},
	{"TN", 24},
	{"TR", 26},
	{"UA", 29},
	{"AE", 23},
	{"GB", 22},
	{"VG", 24}}