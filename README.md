## Microservice for IBAN validation

This microservice is written in Go, without any dependencies. It supports SWIFT rules updates mechanism.

IBAN validation consists of two steps:
1. checking IBAN structure, based on SWIFT rules: https://www.swift.com/standards/data-standards/iban-international-bank-account-number
2. validating IBAN by converting it into an integer and performing a basic mod-97 operation

For more information: https://en.wikipedia.org/wiki/International_Bank_Account_Number#Validating_the_IBAN


###API
Makes a basic GET request with ```iban``` parameter

```curl http://127.0.0.1:8080/?iban=DK5000400440116243```

Response will be returned with ```application/json``` content type:

```json
{
  "iban": "DK5000400440116243",
  "valid": true
}
```
, where
- `iban` (string) - your IBAN
- `valid` (bool) - valid or not

API returns `400 BadRequest` if no `iban` parameter provided:
```json
{
  "error": "iban param required"
}
```


###Makefile
#### ```make build```
Build service binary

#### ```make run```
Run microservice locally

#### ```make test```
Run all tests

#### ```make generate-rules```
Generate new `iban/validator/rules.go` file from SWIFT rules `iban/assets/swift_iban.txt`

#### ```make update-rules```
Download new SWIFT rules file to `iban/assets/swift_iban.txt` and generate `iban/validator/rules.go`

