.PHONY: docs docker docker-down

build:
	go mod tidy
	go build ./cmd/api/.

run:
	./api

docs:
	swag init -g ./cmd/api/main.go --output ./docs

docker:
	docker compose -f ./deployments/docker-compose.yml up -d

docker-down:
	docker compose -f ./deployments/docker-compose.yml  down