# NebengJek
![SonarQube Coverage](https://sonarcloud.io/api/project_badges/measure?project=nebengjek-demo_nebengjek&metric=coverage)
![CI](https://github.com/ermasavior/nebeng-jek/actions/workflows/ci.yml/badge.svg)

## Description
NebengJek is a ride-sharing app that connects users with shared rides. Users can either be Riders, requesting a ride, or Drivers, offering their vehicle. Both can choose whom to ride with, as long as they're within a specific area. 🚀

## Main Features
1. __Ride Matching__

    Riders can request a ride to nearest available drivers (in radius 1 km). Drivers can accept or ignore requests. Riders can also be selective on who offered the ride.
2. __Real-Time Location Tracking__:

    Once the ride is matched and rider got picked up, the ride begins. Both users sending location updates every minute. The app tracks and calculates the distance traveled.
3. __Ride Commissions__:
    
    The app takes 30% commission from each ride to support the service maintenance and growth. :)

## Architecture
![LLD](docs/pictures/Nebengjek-LLD.png)

This system contains of four internal services and one external service (mocked).

1. Riders, responsible for maintaining Riders' connection for real-time location update and ride update broadcast.
2. Drivers, responsible for maintaining Drivers' connection for real-time location update and ride update broadcast.
3. Rides, responsible for managing Ride data, including driver-rider assignments and ride status lifecycle.
4. Location, responsible for managing users' real time locations.
5. (External) Tsel-payment service mock, responsible for maintaining users' credits. 

The communication between users and our services utilize __Websocket__ (for real-time bidirectional communication) and __REST API__ (for stateless data updates). To enable system High Availability, a Load Balancer sits in front of our services. The services are containerized using Docker. The services communicate with Event Driven Architecture using NATS JetStream (for ride update broadcasting and real time location tracking).

The database we are using are Relational Database (Postgres) for storing Rides data and Key-Value storage (Redis) for storing Location data (in Geolocation format). Redis provides geo-location operations, including queries based on indexed coordinates and searches within a specified radius.

## Data Schema
![DB](docs/pictures/ERD.png)

The data consists of four tables: Drivers, Riders, Rides, RideCommissions

### Enumerations
__Ride Status__
![State Diagram](docs/pictures/state-diagram.ride-status.png)

__Driver Status__
| Number | Status    | Description 	              |
|----	 |--------   |-------------	              |
| 0 	 | OFF       | Driver is not in the radar |
| 1  	 | AVAILABLE | Driver turns on the beacon |


## Data Migration
### Prerequisites
1. Postgres 16
2. [DBMate](https://github.com/amacneil/dbmate), data migration script to initialize tables and data seeds

### Steps
Ensure that your database is running. To Initialize schema and add migration:

```
dbmate --url 'postgres://YOUR_USERNAME:YOUR_PASSWORD@DB_HOST:5436/rides_db?sslmode=disable' up
```

## How to Run
### Prerequisites
1. **Golang >=1.22**, serves web API
2. **Postgres 16**, for ride data store
3. **Redis**, for GeoLocation data store
4. **NATS JetStream**, message broker for event streaming and queue group

5. (Alternatively) **Docker**, for practical containerized environment

### Steps

1. Ensure that all the service dependencies are running--Redis, Postgres, and NATS JetStream.
2. Initialize Postgres database (see Data Migration step)
3. Initialize NATS Jetstream using [create_stream script](deployments/nats/create_streams.sh)
2. Initialize `.env` files for each services. (See [`./configs/rides/.env.example`](configs/rides/.env.example) for example)
3. Run each service independently:
    ```sh
    make run-drivers
    make run-riders
    make run-rides
    ```

### Using Docker

1. Initialize `.env` file for each services. [(`./configs/rides/.env.example`)](configs/rides/.env.example)
2. In root path, execute `docker-compose up -d` to run all services (including the dependencies)

## API Contract
TBD

## Load Test

For load testing, we use `k6`.
The test scenario will hit two selected APIs which spawned a number of concurrent users that was run gradually on multiple stages.

The load test target is GET ride data endpoint (with target of 50-200 users) PATCH driver availability endpoint (with target of 20-80 users). 

## Prerequisites
1. [K6](https://github.com/grafana/k6)

## How to Run
To run the load test:
```sh
cd loadtest

## To run load test for GET ride data
k6 run get_ride_data.load_stages.js

## To run load test for PATCH driver availability
k6 run patch_driver_availability.load_stages.js
```

## Author
Erma Safira Nurmasyita

Telegram: @ermasavior
