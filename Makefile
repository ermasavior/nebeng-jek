init:
	chmod +x init-project.sh && ./init-project.sh

run-rides:
	go run cmd/rides/main.go

run-drivers:
	go run cmd/drivers/main.go

run-riders:
	go run cmd/riders/main.go

build:
	go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o bin/rides cmd/rides/main.go

tidy:
	go mod tidy

test:
	go test -v ./...
