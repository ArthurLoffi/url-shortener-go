.PHONY: docs docker docker-down

build:
	go mod tidy
	go build ./cmd/api/.

run:
	./api

docs:
	swag init -g ./cmd/api/main.go --output ./docs

infra:
	docker compose -f ./deployments/docker-compose.yml up -d redis

docker:
	docker compose -f ./deployments/docker-compose.yml up -d

docker-down:
	docker compose -f ./deployments/docker-compose.yml down

docker-rebuild:
	docker compose -f ./deployments/docker-compose.yml down
	docker compose -f ./deployments/docker-compose.yml build --no-cache
	docker compose -f ./deployments/docker-compose.yml up -d