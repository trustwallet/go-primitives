GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin

## generate-coins converts coin/coins.yml file into golang model and helper functions
generate-coins:
	@echo "  >  Generating coin file"
	GOBIN=$(GOBIN) go run -tags=coins coin/gen.go
	goimports -w coin/coins.go

## test executes all tests
test:
	go test -v ./...

## golint: Run linter.
lint: go-lint-install go-lint

go-lint-install:
ifeq (,$(wildcard test -f bin/golangci-lint))
	@echo "  >  Installing golint"
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- v1.54.1
endif

go-lint:
	@echo "  >  Running golint"
	bin/golangci-lint run --timeout=2m
