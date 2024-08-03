all: test build

build:
	go build -o bin/passgen .

test:
	go test -v ./...

clean:
	rm -f bin/passgen

fmt:
	go fmt ./...

.PHONY: all build test clean fmt
