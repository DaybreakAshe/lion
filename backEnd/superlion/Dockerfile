
# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-go-server-ping .

##
## Deploy
##
FROM scratch

WORKDIR /

COPY --from=build /docker-go-server-ping /docker-go-server-ping

EXPOSE 8080

ENTRYPOINT ["/docker-go-server-ping"]

