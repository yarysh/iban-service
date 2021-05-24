GO_CMD=go
GO_APP_PATH=./app
RULES_URL=https://www.swift.com/swift-resource/11971/download

build:
	cd $(GO_APP_PATH) && $(GO_CMD) build -o ../iban-service $$(ls -1 *.go | grep -v _test.go)

run:
	cd $(GO_APP_PATH) && $(GO_CMD) run $$(ls -1 *.go | grep -v _test.go)

test:
	cd $(GO_APP_PATH) && $(GO_CMD) test ./...

generate-rules:
	cd $(GO_APP_PATH) && $(GO_CMD) generate ./iban/validator

update-rules:
	cp -n app/iban/assets/swift_iban.txt{,.bak} 2>/dev/null || :
	curl --request GET -sL --url '$(RULES_URL)' --output 'app/iban/assets/swift_iban.txt'
	make generate-rules
