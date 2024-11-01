run-rides:
	go run cmd/rides/main.go

run-drivers:
	go run cmd/drivers/main.go

run-riders:
	go run cmd/riders/main.go

build-rides:
	cd cmd/rides && env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../../bin/rides main.go

build-drivers:
	cd cmd/drivers && env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../../bin/drivers main.go

build-riders:
	cd cmd/riders && env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../../bin/riders main.go

tidy-rides:
	cd cmd/rides && go mod tidy

tidy-drivers:
	cd cmd/drivers && go mod tidy

tidy-riders:
	cd cmd/riders && go mod tidy

tidy:
	go mod tidy

test:
	go test -v ./...
