APP_NAME := myapp
VERSION ?= $(shell git describe --tags --always)
BUILD_DIR := build
MAIN_PACKAGE := ./server/cmd

.PHONY: all build test lint clean release

all: test lint build

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PACKAGE)

test:
	go test ./...

lint:
	golangci-lint run

clean:
	rm -rf $(BUILD_DIR)

release: clean
	mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-$(VERSION)-linux-amd64 $(MAIN_PACKAGE)
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-$(VERSION)-darwin-amd64 $(MAIN_PACKAGE)
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-$(VERSION)-windows-amd64.exe $(MAIN_PACKAGE)

