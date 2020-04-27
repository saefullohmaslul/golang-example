# Production
FROM golang:1.14-alpine as base
RUN apk update && apk upgrade && \
  apk add --no-cache bash git openssh

RUN mkdir -p /app
WORKDIR /app

# Development
FROM base as dev

RUN apk add --no-cache autoconf automake libtool gettext gettext-dev make g++ texinfo curl
RUN mkdir -p /app

WORKDIR /root

RUN wget https://github.com/emcrisostomo/fswatch/releases/download/1.14.0/fswatch-1.14.0.tar.gz
RUN tar -xvzf fswatch-1.14.0.tar.gz

WORKDIR /root/fswatch-1.14.0
RUN ./configure
RUN make
RUN make install

WORKDIR /app