
# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.18 AS build

WORKDIR /superlion

COPY . ./

ENV GO111MODULE=on

RUN go build -o /lion

##
## Deploy
##
FROM scratch

WORKDIR /

COPY --from=build /lion /lion

EXPOSE 8080

ENTRYPOINT ["/lion"]

