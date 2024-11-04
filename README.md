# Nebengjek

## Prerequisites
1. Docker

## How to Run

1. Initialize `.env` file. See `./configs/.env.example` for example.
2. Run each service independently:
    ```sh
    make run-drivers
    make run-riders
    make run-rides
    ```

### Using Docker

1. Initialize `.env` file into each `./configs/*` folders
2. Go to `/deployments`, then execute `docker-compose up -d` to run dependency services
3. Run docker build and run for each services independently:
    ```sh
    make docker-build-drivers
    make docker-run-drivers

    make docker-build-riders
    make docker-run-riders

    make docker-build-rides
    make docker-run-rides
    ```