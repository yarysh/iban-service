GOCMD=go
RULES_URL=https://www.swift.com/swift-resource/11971/download

build:
	$(GOCMD) build -o iban-service app/*.go

run:
	$(GOCMD) run app/*.go

generate-rules:
	cd app && $(GOCMD) generate ./iban/validator

update-rules:
	cp -n app/iban/assets/swift_iban.txt{,.bak} 2>/dev/null || :
	curl --request GET -sL --url '$(RULES_URL)' --output 'app/iban/assets/swift_iban.txt'
	make generate-rules
