FROM golang:1.17.2-buster as builder
RUN set -xe
RUN apt-get update
RUN apt-get install make gcc g++

WORKDIR /go/src/app_microservice
COPY go.mod go.sum /go/src/app_microservice/

RUN go mod download
