init:
	chmod +x init-project.sh && ./init-project.sh

run:
	go run cmd/api/main.go

build:
	go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o bin/api cmd/api/main.go

tidy:
	go mod tidy

test:
	go test -v ./...
