
# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.18 AS buildimage

WORKDIR /superlion

COPY . ./

ENV GO111MODULE=on

RUN go build -o /lion

##
## Deploy
##
#FROM alpine
#
#WORKDIR /
#
#COPY --from=buildimage /lion ./

EXPOSE 8080

RUN ls -l

ENTRYPOINT ["/lion"]

