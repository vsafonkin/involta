all: build
	
fmt:
	go fmt ./...

lint:
	go vet ./...

build: fmt
	go build -ldflags='-s -w' -o bin/main ./cmd/main.go

run:
	./bin/main

