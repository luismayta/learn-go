# go

GOLANGCI_VERSION ?= 1.39.0

GOPATH	= $(shell go env GOPATH)
GOBIN	= $(GOPATH)/bin

GO_FILES = $(shell find ./ -type f -name '*.go' | grep -v '/vendor/' | sort -u)

# Bin variables
GOLANGCI-LINT = $(GOBIN)/golangci-lint

## show help commands
.PHONY: go.help
go.help:
	@echo '    go:'
	@echo ''
	@echo '        go                 show help'
	@echo '        go.lint            lint go'
	@echo '        go.setup           setup go'
	@echo '        go.fix             fix code of golangci lint'
	@echo '        go.vet             go vet against code'
	@echo '        go.fmt             run fmt for files'
	@echo '        go.build           build application'
	@echo ''

## show help
.PHONY: go
go:
	@if [ -z "${command}" ]; then \
		make go.help;\
	fi

bin/golangci-lint-${GOLANGCI_VERSION}:
	@mkdir -p bin
	[ ! -e "bin/golangci-lint" ] && curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | BINARY=golangci-lint bash -s -- v${GOLANGCI_VERSION}
	@mv bin/golangci-lint $@

bin/golangci-lint: bin/golangci-lint-${GOLANGCI_VERSION}
	@ln -sf golangci-lint-${GOLANGCI_VERSION} bin/golangci-lint

## Run linter go
.PHONY: go.lint
go.lint: bin/golangci-lint
	bin/golangci-lint run --config .github/linters/.golangci.yml

## Fix lint violations
.PHONY: go.fix
go.fix: bin/golangci-lint
	bin/golangci-lint run --fix --config .github/linters/.golangci.yml

## Run go vet against code
.PHONY: go.vet
go.vet:
	go vet ./...

## build go package
.PHONY: go.build
go.build: bin/goreleaser
	bin/goreleaser build --snapshot --rm-dist

## gofmt and goimports all go files
.PHONY: go.fmt
go.fmt:
	gofmt -s -l -w $(GO_FILES)

## setup download and install dependence.
.PHONY: go.setup
go.setup:
	go mod download
	go mod tidy
	go mod vendor
	go generate -v ./...