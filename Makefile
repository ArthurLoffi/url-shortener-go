.PHONY: docs

build:
	go mod tidy
	go build ./cmd/api/.

run:
	./api

docs:
	swag init -g ./cmd/api/main.go --output ./docs