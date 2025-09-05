.PHONY: build clean run up down

build:
	@go build -o bin/server cmd/server/main.go

clean:
	@rm -rf bin

run:
	@go run cmd/server/main.go

up:
	@goose up

down:
	@goose down
