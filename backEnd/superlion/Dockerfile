
# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.18 AS build

WORKDIR /superlion

COPY . ./

ENV GO111MODULE=on

RUN go build -o /docker-go-server-ping

##
## Deploy
##
FROM scratch

WORKDIR /

COPY --from=build /docker-go-server-ping /docker-go-server-ping

EXPOSE 8080

ENTRYPOINT ["/docker-go-server-ping"]

