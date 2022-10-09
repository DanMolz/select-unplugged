.DEFAULT_GOAL := build

lint:
	go vet ./...

test: lint
	go test ./...

build: test
	mkdir -p bin
	GOOS=linux GOARCH=arm GOARM=5 go build -o bin/select-unplugged-linux-arm
	go build -o bin/select-unplugged

cover:
	go test -coverprofile coverage.out ./...
	go tool cover -html=coverage.out