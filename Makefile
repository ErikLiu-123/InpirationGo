.PHONY: dev build run watch

dev:
	air

build:
	go build -o app .

run:
	go run .

watch:
	find . -name "*.go" | entr -r go run .

test:
	go test ./...