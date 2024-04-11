GOOS = $(shell go env GOOS)
GOARCH = $(shell go env GOARCH)
BUILD_DIR = dist/${GOOS}_${GOARCH}

ifeq ($(GOOS),windows)
OUTPUT_PATH = ${BUILD_DIR}/c1-baton-formal.exe
else
OUTPUT_PATH = ${BUILD_DIR}/c1-baton-formal
endif

.PHONY: build
build:
	go build -o ${OUTPUT_PATH} ./cmd/c1-baton-formal
	go build -o ${OUTPUT_PATH} ./cmd/c1-baton-formal

.PHONY: update-deps
update-deps:
	go get -d -u ./...
	go mod tidy -v
	go mod vendor

.PHONY: add-dep
add-dep:
	go mod tidy -v
	go mod vendor

.PHONY: lint
lint:
	golangci-lint run
