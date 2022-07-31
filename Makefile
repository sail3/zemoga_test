default: build

init:
	@cp .env.template .env

build:
	@docker-compose build

devshell:
	@docker-compose run --rm --service-ports service sh

up:
	@docker-compose run --rm --service-ports service

down:
	@docker-compose down --remove-orphans

t:
	@go test -cover ./...

migrate-up:
	migrate -source file:///$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))/migrations -database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=disable up

migrate-down:
	migrate -source file:///$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))/migrations -database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=disable down
