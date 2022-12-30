.DEFAULT_GOAL := build
SHELL := /bin/bash
MAKEFLAGS += --jobs=4

lint:
	go vet ./...

test:
	go test -cover ./...

watch:
	while : ; do \
		make test; \
		date; \
		read -r _; \
	done < <(fswatch -o .)

mkdir-bin:
	mkdir -p bin

build-deps: mkdir-bin lint test

build-arm: build-deps
	GOOS=linux GOARCH=arm GOARM=5 go build -o bin/select-unplugged-linux-arm

build-native: build-deps
	go build -o bin/select-unplugged

build: build-arm build-native

cover:
	go test -coverprofile coverage.out ./...
	go tool cover -html=coverage.out

mocks:
	mockery --inpackage --all
