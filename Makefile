all: lint test

.PHONY: lint
lint:
	golangci-lint run --config .golangci.yml --out-format=tab --tests=false ./...

.PHONY: test
test: lint
	CGO_ENABLED=1 go test -v -race -cover ./...

