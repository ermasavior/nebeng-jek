# NebengJek
![SonarQube Coverage](https://sonarcloud.io/api/project_badges/measure?project=nebengjek-demo_nebengjek&metric=coverage)
![CI](https://github.com/ermasavior/nebeng-jek/actions/workflows/ci.yml/badge.svg)

## Description
NebengJek is a ride-sharing app that connects users with shared rides. Users can either be Riders, requesting a ride, or Drivers, offering their vehicle. Both can choose whom to ride with, as long as they're within a specific area. 🚀

## Features
1. __Ride Matchmaking__

    Riders can request a rides to nearest available drivers (in radius 1 km). Drivers can accept or ignore requests. Riders can also be selective on who offered the ride.
2. __Real-Time Location Tracking__:

    Once the ride is matched and rider got picked up, the ride begins, with both users sending location updates every minute. The app tracks and calculates the distance traveled.
3. __Ride Commissions__:
    
    The app takes a 30% commission from each ride to support the service maintenance and growth. :)

## Architecture
### High Level Design
![HLD](docs/pictures/Nebengjek-HLD.png)

__Services__

1. Riders, responsible for maintaining Riders' connection for real-time location update and ride update broadcast.
2. Drivers, responsible for maintaining Drivers' connection for real-time location update and ride update broadcast.
3. Rides, responsible for managing Ride data, including driver-rider assignments and ride status lifecycle.
4. Location, responsible for managing users' real time locations.
5. (External) Tsel-payment service mock, responsible for maintaining users' credit. 

### Low Level Design
![LLD](docs/pictures/Nebengjek-LLD.png)

The communication between users and our services utilize __Websocket__ (for real time bidirectional communication) and __REST API__ (for stateless data update). For distributing user traffic, we use Load Balancer in front of the services. The services are containerized using Docker. The services communicate with pubsub using NATS message broker (for ride update broadcasting and real time location).

The database we are using are Relational Database (Postgres) for storing Rides data and Key-Value storage (Redis) for storing Location data (in Geolocation format).

## Data Schema
![DB](docs/pictures/ERD.png)

TBD
### Drivers
### Riders
### Rides
### Ride Commissions

<!-- ### Ride status
- NEW_RIDE_REQUEST
- MATCHED_DRIVER
- READY_TO_PICKUP
- RIDE_STARTED
- RIDE_ENDED
- RIDE_PAID
- RIDE_CANCELLED -->

## Data Migration
### Prerequisites
1. Postgres 16
2. DBMate

### Steps
Ensure that your database is running. To Initialize schema and add migration:

```
dbmate --url 'postgres://rides_username:rides_password@0127.0.0.1:5436/rides_db?sslmode=disable' up
```

## How to Run
### Prerequisites
1. Golang >=1.22
2. Postgres
3. DBMate (data migration tools)

### Steps

1. Ensure that all the service dependencies are active: Redis, Postgres, and NATS. (You can use the dockerized version in `docker-compose.yml`)
2. Initialize `.env` file. See `./configs/.env.example` for example.
3. Run each service independently:
    ```sh
    make run-drivers
    make run-riders
    make run-rides
    ```

### Using Docker

1. Initialize `.env` file into each `./configs/*` folders
2. In root path, execute `docker-compose up -d` to run all services (including the dependencies)

## API Contract
TBD

## Load Test

For load testing, we use `k6`.
The test scenario will hit two selected APIs which spawned a number of concurrent users that was run gradually on multiple stages.

The load test target is GET ride data endpoint (with target of 50-200 users) PATCH set driver availability endpoint (with target of 20-80 users). 

## How to Run
To run the load test:
```sh
cd loadtest && k6 run load_stages.js 
```

## Author
Erma Safira Nurmasyita

Telegram: @ermasavior
