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
2. Go to `/deployments`, then execute `docker-compose up -d` to run all services (including the dependencies)

## Load Test

For load testing, we use `k6`.
The load test target is PATCH set driver availability endpoint. The test scenario spawned 50-250 concurrent users that run gradually on stages.

To run the load test:
```sh
cd loadtest && k6 run load_stages.js 
```