GO_VERSION_SHORT:=$(shell echo `go version` | sed -E 's/.* go(.*) .*/\1/g')
ifneq ("1.16","$(shell printf "$(GO_VERSION_SHORT)\n1.16" | sort -V | head -1)")
$(error NEED GO VERSION >= 1.16. Found: $(GO_VERSION_SHORT))
endif

export GO111MODULE=on

SERVICE_NAME=omp-bot
SERVICE_PATH=Valeviy/omp-bot

PGV_VERSION:="v0.6.1"

OS_NAME=$(shell uname -s)
OS_ARCH=$(shell uname -m)
GO_BIN=$(shell go env GOPATH)/bin

ifeq ("NT", "$(findstring NT,$(OS_NAME))")
OS_NAME=Windows
endif

.PHONY: run
run:
	go run cmd/omp-bot/main.go

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	go test -v -race -timeout 30s -coverprofile cover.out ./...
	go tool cover -func cover.out | grep total | awk '{print $$3}'
# ----------------------------------------------------------------
.PHONY: build
build: .build

.build:
	go mod download && CGO_ENABLED=0  go build \
		-tags='no_mysql no_sqlite3' \
		-ldflags=" \
			-X 'github.com/$(SERVICE_PATH)/internal/config.version=$(VERSION)' \
			-X 'github.com/$(SERVICE_PATH)/internal/config.commitHash=$(COMMIT_HASH)' \
		" \
		-o ./bin/omp-bot$(shell go env GOEXE) ./cmd/omp-bot/main.go
