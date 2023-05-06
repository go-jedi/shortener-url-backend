.PHONY: run

build:
	go build -o ./.bin/shortener-url cmd/shortener-url/main.go

run: build
	./.bin/shortener-url

.DEFAULT_GOAL := run 