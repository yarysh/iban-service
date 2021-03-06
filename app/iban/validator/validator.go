package validator

const MOD = 97

//go:generate go run "github.com/yarysh/iban-service/app/iban" gen_validator_rules
//Validate - check if IBAN is valid
// Validation consists of two steps:
//	1. checking IBAN structure, based on SWIFT rules: https://www.swift.com/standards/data-standards/iban-international-bank-account-number
//	2. validating IBAN by converting it into an integer and performing a basic mod-97 operation
// More information: https://en.wikipedia.org/wiki/International_Bank_Account_Number#Validating_the_IBAN
func Validate(iban string) bool {
	ibanLen := len(iban)
	if ibanLen < 2 {
		return false
	}

	country := iban[0:2]
	rule, ok := Rules[country]
	if !ok || ibanLen != rule.Size {
		return false
	}

	if !rule.Regex.MatchString(iban) {
		return false
	}

	total := 0
	pos, i := 4, 0
	for ; i < ibanLen; i++ {
		if pos == ibanLen {
			pos = 0
		}
		char := rune(iban[pos])
		var n int
		if '0' <= char && char <= '9' {
			n = int(char) - '0'
		} else if 'A' <= char && char <= 'Z' {
			n = 10 + int(char) - 'A'
		}
		if n > 9 {
			total = (100*total + n) % MOD
		} else {
			total = (10*total + n) % MOD
		}
		pos++
	}
	return total == 1
}
