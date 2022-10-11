.DEFAULT_GOAL := build
SHELL := /bin/bash

lint:
	go vet ./...

test: lint
	go test -cover ./...

watch:
	while : ; do \
		make test; \
		date; \
		read -r _; \
	done < <(fswatch -o .)

build: test
	mkdir -p bin
	GOOS=linux GOARCH=arm GOARM=5 go build -o bin/select-unplugged-linux-arm
	go build -o bin/select-unplugged

cover:
	go test -coverprofile coverage.out ./...
	go tool cover -html=coverage.out

mocks:
	mockery --inpackage --all
