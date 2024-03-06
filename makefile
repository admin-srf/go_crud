.PHONY: default run test build clean

APP_NAME=go_crud

default: run-with-docs


dependencies:
	@go install

test:
	@go test

run:
	@go run main.go

run-with-docs:
	@swag init --parseDependency --parseInternal
	@go run main.go

build:
	@go build -o $(APP_NAME) main.go


docs:
	@swag init