#!/bin/sh

# Create Ride Events Stream
nats stream add RIDE_EVENTS \
    --server nats://localhost:4222 \
    --subjects "ride.*" \
    --storage file \
    --retention limits \
    --discard old \
    --max-msgs 10000 \
    --replicas 1 \
    --max-age 1h \
    --max-bytes 5242880 \
    --max-msg-size 1048576 \
    --dupe-window 2m \
    --ack

# Create User Live Location Stream
nats stream add USER_LIVE_LOCATION \
    --server nats://localhost:4222 \
    --subjects "user.live_location" \
    --storage file \
    --retention limits \
    --discard old \
    --max-msgs 10000 \
    --replicas 1 \
    --max-age 1h \
    --max-bytes 5242880 \
    --max-msg-size 1048576 \
    --dupe-window 2m \
    --ack

nats consumer add RIDE_EVENTS riders_service --defaults --ack=explicit --target=pull --wait=3m --max-deliver 5
nats consumer add RIDE_EVENTS drivers_service --defaults --ack=explicit --target=pull --wait=3m --max-deliver 5
nats consumer add USER_LIVE_LOCATION rides_service --defaults --ack=explicit --target=pull --wait=3m --max-deliver 5
