run-rides:
	ENV_PATH=./configs/rides/.env go run cmd/rides/main.go

run-drivers:
	ENV_PATH=./configs/drivers/.env go run cmd/drivers/main.go

run-riders:
	ENV_PATH=./configs/riders/.env go run cmd/riders/main.go

run-location:
	ENV_PATH=./configs/location/.env go run cmd/location/main.go


build-rides:
	cd cmd/rides && env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../../bin/rides main.go

build-drivers:
	cd cmd/drivers && env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../../bin/drivers main.go

build-riders:
	cd cmd/riders && env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../../bin/riders main.go

build-location:
	cd cmd/location && env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../../bin/location main.go


docker-run-drivers:
	docker build -f deployments/drivers/Dockerfile -t nebengjek/drivers:v1.0.0 .
	docker run --env-file ./configs/drivers/.env --rm -p 9998:9998 --network deployments_local --name nebengjek-drivers nebengjek/drivers:v1.0.0

docker-run-riders:
	docker build -f deployments/riders/Dockerfile -t nebengjek/riders:v1.0.0 .
	docker run --env-file ./configs/riders/.env --rm -p 9997:9997 --network deployments_local --name nebengjek-riders nebengjek/riders:v1.0.0

docker-run-rides:
	docker build -f deployments/rides/Dockerfile -t nebengjek/rides:v1.0.0 .
	docker run --env-file ./configs/rides/.env --rm -p 9999:9999 --network deployments_local --name nebengjek-rides nebengjek/rides:v1.0.0

docker-run-location:
	docker build -f deployments/location/Dockerfile -t nebengjek/location:v1.0.0 .
	docker run --env-file ./configs/location/.env --rm -p 9996:9996 --network deployments_local --name nebengjek-location nebengjek/location:v1.0.0


tidy-rides:
	cd cmd/rides && go mod tidy

tidy-drivers:
	cd cmd/drivers && go mod tidy

tidy-riders:
	cd cmd/riders && go mod tidy

tidy-location:
	cd cmd/location && go mod tidy


tidy:
	go mod tidy

test:
	go test -v ./...
