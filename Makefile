TAG ?= latest
IMG ?= level
REGISTRY := dukeman

fmt:
	go fmt ./...

test:
	go test ./... -race

lint:
	golangci-lint run

docker-build:
	DOCKER_BUILDKIT=1 docker build --progress=plain -t $(REGISTRY)/$(IMG):$(TAG) .