package validator

import (
	"fmt"
)

//go:generate go run "github.com/yarysh/iban-service/app/iban" gen_validator_rules
func Validate(iban string) bool {
	for _, rule := range Rules {
		fmt.Println(rule)
	}
	return false
}
