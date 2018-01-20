FROM node:6-alpine

MAINTAINER evalsocket<evalsocket@protonmail.com>

RUN cd $HOME
RUN mkdir -p go/src/github.com/bakins
RUN cd go/src/github.com/bakins
RUN git clone github.com/bakins/grpc-fastcgi-proxy
RUN cd grpc-fastcgi-proxy

RUN go build ./cmd/grpc-fastcgi-proxy
