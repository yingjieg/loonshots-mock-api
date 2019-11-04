.PHONY: build run reflex

test:
	go test ./tests -v

build:
	go build -o ./build/mock-api ./server.go

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./build/mock-api-linux ./server.go

run: build
	./build/mock-api

reflex:
	reflex --start-service -r '\.go$$' make run
