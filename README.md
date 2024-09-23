# Telkomsel Go Boilerplate

The official golang boilerplate in Telkomsel. This boilerplate consists of a simple REST web server with create, read, update, delete (CRUD) functionalities.

## Library

1. Web Framework: Gin
2. Config: Godotenv & Consul API
3. Logger: Zap
4. Monitoring: OpenTelemetry SDK
5. Database: SQLx
6. Testing tools: SQLMock, Testify

## How to Run

1. Run `make init` script to begin renaming the project name. Then, input your project name.
2. Import db schema in data/postgres.sql to a dummy postgres db.
3. Initialize `configs/.env` file
4. Execute:
```
make run
```

### Using Docker

1. Initialize `configs/.env` file
2. Run `docker-compose up postgres -d` or `docker-compose up redis -d` to run services separately
3. To run the service:
    ```sh
    docker build -f ./Dockerfile -t boilerplate-app .
    docker run --name boilerplate-app -p 9999:9999 --env-file ./configs/.env --network deployments_local boilerplate-app
    ```