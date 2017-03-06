.PHONY: all

all: clean build

install:
	curl https://glide.sh/get | sh
	glide install

build:
	go build -o bin/server ./cmd/server

test:
	go test $(shell go list ./... | grep -Ev 'vendor')

clean:
	rm -f bin/*