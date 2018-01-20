FROM golang:alpine

MAINTAINER evalsocket<evalsocket@protonmail.com>

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN apk add lighttpd php5-common php5-iconv php5-json php5-gd php5-curl php5-xml php5-pgsql php5-imap php5-cgi fcgi

RUN # Environments
ENV TIMEZONE            Asia/Jakarta
ENV PHP_MEMORY_LIMIT    510M
ENV MAX_UPLOAD          50M
ENV PHP_MAX_FILE_UPLOAD 200
ENV PHP_MAX_POST        100M                                                                                                                               

RUN mkdir -p /go/src/github.com/bakins
WORKDIR /go/src/github.com/bakins
RUN git clone https://github.com/evalsocket/grpc-fastcgi-proxy.git
WORKDIR ./grpc-fastcgi-proxy
RUN pwd
RUN go build ./cmd/grpc-fastcgi-proxy

EXPOSE 8080

RUN ./grpc-fastcgi-proxy $HOME/git/grpc-fastcgi-example/index.php


