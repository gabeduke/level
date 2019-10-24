# Docker config
TAG ?= latest
IMG ?= level
REGISTRY ?= dukeman
DOCKER_IMG = $(REGISTRY)/$(IMG):$(TAG)
PORT = 8080

.DEFAULT_GOAL = help

.PHONY: help fmt test lint docker-build

fmt: ## fmt project
	go fmt ./...

test: ## test project
	go test ./... -race -coverprofile=coverage.txt -covermode=atomic

lint: ## lint project
	golangci-lint run

swagger: swagger-init swagger-static swagger-readme ## rebuild swagger docs

swagger-init:
	swag init

swagger-static:
	docker run --rm -v ${PWD}:/local --user $(shell id -u):$(shell id -u) \
		swaggerapi/swagger-codegen-cli generate \
		-i /local/docs/swagger.yaml \
		-l html2 \
		-o /local/docs

swagger-readme:
	docker run --rm \
		--user $(shell id -u):$(shell id -u) \
		--volume $(shell pwd):/app \
		--workdir /app \
			node npx markdown-swagger /app/docs/swagger.yaml /app/README.md

build: swagger ## build container
	DOCKER_BUILDKIT=1 docker build -t $(DOCKER_IMG) .

dev: swagger ## run program in dev mode
	go run main.go

run: docker ## run project in container
	docker run -p $(PORT):8080 -it $(DOCKER_IMG)

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'