include $(CURDIR)/bootstrap/Makefile
SERVICE_NAME := level
VAR_FILE := $(CURDIR)/level.tfvars

# Docker config
TAG ?= latest
IMG ?= level
REGISTRY ?= dukeman
DOCKER_IMG = $(REGISTRY)/$(IMG):$(TAG)
PORT = 8080
MODULE = github.com/gabeduke/level

.DEFAULT_GOAL = help

##########################################################
##@ APP
##########################################################
.PHONY: build dev run

build: swagger											## build container
	DOCKER_BUILDKIT=1 docker build -t $(DOCKER_IMG) .

dev: swagger											## run program in dev mode
	$(info INFO serving API in 'dev' mode)
	LOG_LEVEL=debug go run main.go

run: docker												## run project in container
	docker run -p $(PORT):8080 -it $(DOCKER_IMG)

##########################################################
##@ TEST
##########################################################
.PHONY: fmt test lint

fmt:													## fmt project
	go fmt ./...

test:													## test project
	go test ./... -race -coverprofile=coverage.txt -covermode=atomic

lint:													## lint project
	golangci-lint run

##########################################################
##@ DOCS
##########################################################
.PHONY: docs swagger swagger-init swagger-static swagger-readme

docs:													## Serve package godocs
	$(info http://localhost:6060/pkg/$(MODULE))
	godoc -http=localhost:6060

swagger: swagger-init swagger-static swagger-readme		## rebuild swagger docs

swagger-init:
	$(info INFO init swagger)
	@swag init

swagger-static:
	$(info INFO generating swagger static files)
	@docker run --rm -v ${PWD}:/local --user $(shell id -u):$(shell id -u) \
		swaggerapi/swagger-codegen-cli generate \
		-i /local/docs/swagger.yaml \
		-l html2 \
		-o /local/docs

swagger-readme:
	$(info INFO generating swagger readme)
	@docker run --rm \
		--user $(shell id -u):$(shell id -u) \
		--volume $(shell pwd):/app \
		--workdir /app \
			node npx markdown-swagger /app/docs/swagger.yaml /app/README.md

##########################################################
##@ UTIL
##########################################################
.PHONY: help

help:													## show help
		@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m 	%s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
