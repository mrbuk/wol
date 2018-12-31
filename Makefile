name = wol

all: build
.PHONY : all

test:
	go test ./...

lint:
	golint .

build: test lint
	mkdir -p build
	go build -o ./build/${name} cmd/wol.go
