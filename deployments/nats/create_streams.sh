#!/bin/sh

# Create Ride Events Stream
nats stream add RIDE_EVENTS \
    --server nats://nats_container:4222 \
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
    --no-ack

# Create User Live Location Stream
nats stream add USER_LIVE_LOCATION \
    --server nats://nats_container:4222 \
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
    --no-ack