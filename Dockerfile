FROM golang:alpine

MAINTAINER evalsocket<evalsocket@protonmail.com>

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN apk add lighttpd php5-common php5-iconv php5-json php5-gd php5-curl php5-xml php5-pgsql php5-imap php5-cgi fcgi
                                                                                                                        
RUN mkdir -p /go/src/github.com/bakins
WORKDIR /go/src/github.com/bakins
COPY . /go/src/github.com/bakins
WORKDIR ./grpc-fastcgi-proxy
RUN pwd
RUN go build ./cmd/grpc-fastcgi-proxy

EXPOSE 8080

RUN 













