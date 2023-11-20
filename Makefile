all: generate build

generate:
	go generate ./...

test:
	go test ./...

build:
	go build -o bin/ballistic ./ballistic
	go build -o bin/proxy ./proxy
