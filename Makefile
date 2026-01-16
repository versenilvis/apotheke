.PHONY: build install clean test

build:
	@mkdir -p bin
	@go build -o bin/apotheke ./cmd

install: build
	@cp bin/apotheke ~/.local/bin/

clean:
	@rm -rf bin

test:
	@go test ./...
