FROM golang:1.22.4 AS golang

WORKDIR /app
COPY ./.. .
RUN make build-location

FROM alpine:3.18.2 AS alpine
RUN apk update && \
    apk add --no-cache ca-certificates=20240226-r0 tzdata=2024a-r0 && update-ca-certificates

FROM alpine:3.18.2

WORKDIR /app
COPY --from=alpine /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=alpine /etc/passwd /etc/passwd
COPY --from=golang /app/bin/location /app/location

ENTRYPOINT ["./location"]