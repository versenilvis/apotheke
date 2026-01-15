.PHONY: build install clean test

build:
	@go build -o apotheke ./cmd/apotheke

install: build
	@sudo cp apotheke /usr/local/bin/

clean:
	@rm -f apotheke

test:
	@go test ./...
