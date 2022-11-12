FROM golang:1.19 AS build-env
WORKDIR /go/src/

COPY . .

RUN CGO_ENABLED=0 \
    GOOS=linux \
    go build \
    -trimpath \
    -o bin/run ./backend/go/cmd/*.go

FROM alpine:latest AS cert
RUN apk --update add ca-certificates
RUN apk --no-cache add tzdata

FROM scratch
WORKDIR /

COPY --from=build-env /go/src/bin/run .
COPY --from=cert /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=cert /usr/share/zoneinfo /usr/share/zoneinfo

ENTRYPOINT ["./run"]