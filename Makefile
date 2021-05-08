.PHONY: clean

all: build

build:
	go mod download
	go build --ldflags "-s -w" -o bin/comet ./cmd/comet/main.go

format:
	go fmt ./...

clean:
	rm -rf bin/
