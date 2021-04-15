.PHONY: deps build test mock cloc unit-test testonly


deps:
	env GO111MODULE=on go mod download
	env GO111MODULE=on go mod vendor


build: deps 

lint:
	golangci-lint run

test: lint deps unit-test

testonly: deps unit-test

cloc:
	cloc --exclude-dir=vendor,3rdmocks,mocks,tools --not-match-f=test .

unit-test:
	go vet `go list ./... | grep -v '/vendor/' | grep -v '/tools'`
	go test -count=1 -cover ./...
