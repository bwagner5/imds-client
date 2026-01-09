BUILD_DIR ?= $(dir $(realpath -s $(firstword $(MAKEFILE_LIST))))/build
VERSION ?= $(shell git describe --tags --always --dirty)
GOOS ?= $(shell uname | tr '[:upper:]' '[:lower:]')
GOARCH ?= $(shell [[ `uname -m` = "x86_64" ]] && echo "amd64" || echo "arm64" )
GOPROXY ?= "https://proxy.golang.org|direct"

$(shell mkdir -p ${BUILD_DIR})

all: fmt verify test build

codegen: ## Generate the IMDS SDK
	go run codegen/staticmetadata.go > pkg/imds/zz_metadata.go
	go run codegen/docs/docs.go > pkg/docs/zz_docs.go

build: ## build binary using current OS and Arch
	go build -a -ldflags="-s -w -X main.version=${VERSION}" -o ${BUILD_DIR}/imds-${GOOS}-${GOARCH} ${BUILD_DIR}/../cmd/main.go

build-local: ## build binary for local development
	go build -ldflags="-X main.version=${VERSION}" -o imds ./cmd/main.go

test: ## run go tests and benchmarks
	go test -bench=. ${BUILD_DIR}/../... -v -coverprofile=coverage.out -covermode=atomic -outputdir=${BUILD_DIR}

version: ## Output version of local HEAD
	@echo ${VERSION}

verify: ## Run Verifications like helm-lint and govulncheck
	go run golang.org/x/vuln/cmd/govulncheck ./...
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run

fmt: ## go fmt the code
	find . -iname "*.go" -exec go fmt {} \;

toolchain: ## Install the development toolchain
	go mod download

licenses: ## Verifies dependency licenses
	go mod download
	! go run github.com/google/go-licenses csv ./... | grep -v -e 'MIT' -e 'Apache-2.0' -e 'BSD-3-Clause' -e 'BSD-2-Clause' -e 'ISC' -e 'MPL-2.0'

help: ## Display help
	@awk 'BEGIN {FS = ":.*##"; printf "Usage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: all build test verify help codegen licenses fmt version toolchain
