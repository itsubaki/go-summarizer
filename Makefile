.PHONY: setup build install lint

setup:
	go mod tidy

run: setup
	go run . $(URL)

build:
	go build .

install:
	go install .

lint:
	golangci-lint run
