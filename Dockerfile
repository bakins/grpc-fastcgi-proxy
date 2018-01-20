FROM golang:alpine

MAINTAINER evalsocket<evalsocket@protonmail.com>

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

RUN mkdir -p /go/src/github.com/bakins
WORKDIR /go/src/github.com/bakins
RUN git clone https://github.com/bakins/grpc-fastcgi-proxy.git
WORKDIR ./grpc-fastcgi-proxy
RUN pwd
RUN go build ./cmd/grpc-fastcgi-proxy
RUN ./grpc-fastcgi-proxy $HOME/git/grpc-fastcgi-example/index.php


