FROM golang:1.22.4 as golang

WORKDIR /app
COPY ./.. .
RUN make build-drivers

FROM alpine:3.18.2 as alpine
RUN apk update && apk add --no-cache ca-certificates tzdata && update-ca-certificates

FROM alpine:3.18.2

WORKDIR /app
COPY --from=alpine /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=alpine /etc/passwd /etc/passwd
COPY --from=golang /app/bin/drivers /app/drivers

ENTRYPOINT ["./drivers"]