// Code generated by go generate; DO NOT EDIT
// Use go generate ./iban/validator

package validator

import "regexp"

type Rule struct {
	Regex *regexp.Regexp
	Size  int
}

var Rules = map[string]Rule{
	"AD": {regexp.MustCompile("^AD[0-9]{2}[0-9]{4}[0-9]{4}[A-z0-9]{12}$"), 24},
	"AE": {regexp.MustCompile("^AE[0-9]{2}[0-9]{3}[0-9]{16}$"), 23},
	"AL": {regexp.MustCompile("^AL[0-9]{2}[0-9]{8}[A-z0-9]{16}$"), 28},
	"AT": {regexp.MustCompile("^AT[0-9]{2}[0-9]{5}[0-9]{11}$"), 20},
	"AZ": {regexp.MustCompile("^AZ[0-9]{2}[A-Z]{4}[A-z0-9]{20}$"), 28},
	"BA": {regexp.MustCompile("^BA[0-9]{2}[0-9]{3}[0-9]{3}[0-9]{8}[0-9]{2}$"), 20},
	"BE": {regexp.MustCompile("^BE[0-9]{2}[0-9]{3}[0-9]{7}[0-9]{2}$"), 16},
	"BG": {regexp.MustCompile("^BG[0-9]{2}[A-Z]{4}[0-9]{4}[0-9]{2}[A-z0-9]{8}$"), 22},
	"BH": {regexp.MustCompile("^BH[0-9]{2}[A-Z]{4}[A-z0-9]{14}$"), 22},
	"BR": {regexp.MustCompile("^BR[0-9]{2}[0-9]{8}[0-9]{5}[0-9]{10}[A-Z][A-z0-9]$"), 29},
	"BY": {regexp.MustCompile("^BY[0-9]{2}[A-z0-9]{4}[0-9]{4}[A-z0-9]{16}$"), 28},
	"CH": {regexp.MustCompile("^CH[0-9]{2}[0-9]{5}[A-z0-9]{12}$"), 21},
	"CR": {regexp.MustCompile("^CR[0-9]{2}[0-9]{4}[0-9]{14}$"), 22},
	"CY": {regexp.MustCompile("^CY[0-9]{2}[0-9]{3}[0-9]{5}[A-z0-9]{16}$"), 28},
	"CZ": {regexp.MustCompile("^CZ[0-9]{2}[0-9]{4}[0-9]{6}[0-9]{10}$"), 24},
	"DE": {regexp.MustCompile("^DE[0-9]{2}[0-9]{8}[0-9]{10}$"), 22},
	"DK": {regexp.MustCompile("^DK[0-9]{2}[0-9]{4}[0-9]{9}[0-9]$"), 18},
	"DO": {regexp.MustCompile("^DO[0-9]{2}[A-z0-9]{4}[0-9]{20}$"), 28},
	"EE": {regexp.MustCompile("^EE[0-9]{2}[0-9]{2}[0-9]{2}[0-9]{11}[0-9]$"), 20},
	"EG": {regexp.MustCompile("^EG[0-9]{2}[0-9]{4}[0-9]{4}[0-9]{17}$"), 29},
	"ES": {regexp.MustCompile("^ES[0-9]{2}[0-9]{4}[0-9]{4}[0-9][0-9][0-9]{10}$"), 24},
	"FI": {regexp.MustCompile("^FI[0-9]{2}[0-9]{3}[0-9]{11}$"), 18},
	"FO": {regexp.MustCompile("^FO[0-9]{2}[0-9]{4}[0-9]{9}[0-9]$"), 18},
	"FR": {regexp.MustCompile("^FR[0-9]{2}[0-9]{5}[0-9]{5}[A-z0-9]{11}[0-9]{2}$"), 27},
	"GB": {regexp.MustCompile("^GB[0-9]{2}[A-Z]{4}[0-9]{6}[0-9]{8}$"), 22},
	"GE": {regexp.MustCompile("^GE[0-9]{2}[A-Z]{2}[0-9]{16}$"), 22},
	"GI": {regexp.MustCompile("^GI[0-9]{2}[A-Z]{4}[A-z0-9]{15}$"), 23},
	"GL": {regexp.MustCompile("^GL[0-9]{2}[0-9]{4}[0-9]{9}[0-9]$"), 18},
	"GR": {regexp.MustCompile("^GR[0-9]{2}[0-9]{3}[0-9]{4}[A-z0-9]{16}$"), 27},
	"GT": {regexp.MustCompile("^GT[0-9]{2}[A-z0-9]{4}[A-z0-9]{20}$"), 28},
	"HR": {regexp.MustCompile("^HR[0-9]{2}[0-9]{7}[0-9]{10}$"), 21},
	"HU": {regexp.MustCompile("^HU[0-9]{2}[0-9]{3}[0-9]{4}[0-9][0-9]{15}[0-9]$"), 28},
	"IE": {regexp.MustCompile("^IE[0-9]{2}[A-Z]{4}[0-9]{6}[0-9]{8}$"), 22},
	"IL": {regexp.MustCompile("^IL[0-9]{2}[0-9]{3}[0-9]{3}[0-9]{13}$"), 23},
	"IQ": {regexp.MustCompile("^IQ[0-9]{2}[A-Z]{4}[0-9]{3}[0-9]{12}$"), 23},
	"IS": {regexp.MustCompile("^IS[0-9]{2}[0-9]{4}[0-9]{2}[0-9]{6}[0-9]{10}$"), 26},
	"IT": {regexp.MustCompile("^IT[0-9]{2}[A-Z][0-9]{5}[0-9]{5}[A-z0-9]{12}$"), 27},
	"JO": {regexp.MustCompile("^JO[0-9]{2}[A-Z]{4}[0-9]{4}[A-z0-9]{18}$"), 30},
	"KW": {regexp.MustCompile("^KW[0-9]{2}[A-Z]{4}[A-z0-9]{22}$"), 30},
	"KZ": {regexp.MustCompile("^KZ[0-9]{2}[0-9]{3}[A-z0-9]{13}$"), 20},
	"LB": {regexp.MustCompile("^LB[0-9]{2}[0-9]{4}[A-z0-9]{20}$"), 28},
	"LC": {regexp.MustCompile("^LC[0-9]{2}[A-Z]{4}[A-z0-9]{24}$"), 32},
	"LI": {regexp.MustCompile("^LI[0-9]{2}[0-9]{5}[A-z0-9]{12}$"), 21},
	"LT": {regexp.MustCompile("^LT[0-9]{2}[0-9]{5}[0-9]{11}$"), 20},
	"LU": {regexp.MustCompile("^LU[0-9]{2}[0-9]{3}[A-z0-9]{13}$"), 20},
	"LV": {regexp.MustCompile("^LV[0-9]{2}[A-Z]{4}[A-z0-9]{13}$"), 21},
	"LY": {regexp.MustCompile("^LY[0-9]{2}[0-9]{3}[0-9]{3}[0-9]{15}$"), 25},
	"MC": {regexp.MustCompile("^MC[0-9]{2}[0-9]{5}[0-9]{5}[A-z0-9]{11}[0-9]{2}$"), 27},
	"MD": {regexp.MustCompile("^MD[0-9]{2}[A-z0-9]{2}[A-z0-9]{18}$"), 24},
	"ME": {regexp.MustCompile("^ME[0-9]{2}[0-9]{3}[0-9]{13}[0-9]{2}$"), 22},
	"MK": {regexp.MustCompile("^MK[0-9]{2}[0-9]{3}[A-z0-9]{10}[0-9]{2}$"), 19},
	"MR": {regexp.MustCompile("^MR[0-9]{2}[0-9]{5}[0-9]{5}[0-9]{11}[0-9]{2}$"), 27},
	"MT": {regexp.MustCompile("^MT[0-9]{2}[A-Z]{4}[0-9]{5}[A-z0-9]{18}$"), 31},
	"MU": {regexp.MustCompile("^MU[0-9]{2}[A-Z]{4}[0-9]{2}[0-9]{2}[0-9]{12}[0-9]{3}[A-Z]{3}$"), 30},
	"NL": {regexp.MustCompile("^NL[0-9]{2}[A-Z]{4}[0-9]{10}$"), 18},
	"NO": {regexp.MustCompile("^NO[0-9]{2}[0-9]{4}[0-9]{6}[0-9]$"), 15},
	"PK": {regexp.MustCompile("^PK[0-9]{2}[A-Z]{4}[A-z0-9]{16}$"), 24},
	"PL": {regexp.MustCompile("^PL[0-9]{2}[0-9]{8}[0-9]{16}$"), 28},
	"PS": {regexp.MustCompile("^PS[0-9]{2}[A-Z]{4}[A-z0-9]{21}$"), 29},
	"PT": {regexp.MustCompile("^PT[0-9]{2}[0-9]{4}[0-9]{4}[0-9]{11}[0-9]{2}$"), 25},
	"QA": {regexp.MustCompile("^QA[0-9]{2}[A-Z]{4}[A-z0-9]{21}$"), 29},
	"RO": {regexp.MustCompile("^RO[0-9]{2}[A-Z]{4}[A-z0-9]{16}$"), 24},
	"RS": {regexp.MustCompile("^RS[0-9]{2}[0-9]{3}[0-9]{13}[0-9]{2}$"), 22},
	"SA": {regexp.MustCompile("^SA[0-9]{2}[0-9]{2}[A-z0-9]{18}$"), 24},
	"SC": {regexp.MustCompile("^SC[0-9]{2}[A-Z]{4}[0-9]{2}[0-9]{2}[0-9]{16}[A-Z]{3}$"), 31},
	"SE": {regexp.MustCompile("^SE[0-9]{2}[0-9]{3}[0-9]{16}[0-9]$"), 24},
	"SI": {regexp.MustCompile("^SI[0-9]{2}[0-9]{5}[0-9]{8}[0-9]{2}$"), 19},
	"SK": {regexp.MustCompile("^SK[0-9]{2}[0-9]{4}[0-9]{6}[0-9]{10}$"), 24},
	"SM": {regexp.MustCompile("^SM[0-9]{2}[A-Z][0-9]{5}[0-9]{5}[A-z0-9]{12}$"), 27},
	"ST": {regexp.MustCompile("^ST[0-9]{2}[0-9]{4}[0-9]{4}[0-9]{11}[0-9]{2}$"), 25},
	"SV": {regexp.MustCompile("^SV[0-9]{2}[A-Z]{4}[0-9]{20}$"), 28},
	"TL": {regexp.MustCompile("^TL[0-9]{2}[0-9]{3}[0-9]{14}[0-9]{2}$"), 23},
	"TN": {regexp.MustCompile("^TN[0-9]{2}[0-9]{2}[0-9]{3}[0-9]{13}[0-9]{2}$"), 24},
	"TR": {regexp.MustCompile("^TR[0-9]{2}[0-9]{5}[0-9][A-z0-9]{16}$"), 26},
	"UA": {regexp.MustCompile("^UA[0-9]{2}[0-9]{6}[A-z0-9]{19}$"), 29},
	"VA": {regexp.MustCompile("^VA[0-9]{2}[0-9]{3}[0-9]{15}$"), 22},
	"VG": {regexp.MustCompile("^VG[0-9]{2}[A-Z]{4}[0-9]{16}$"), 24},
	"XK": {regexp.MustCompile("^XK[0-9]{2}[0-9]{4}[0-9]{10}[0-9]{2}$"), 20},
}
