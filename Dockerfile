# Production
FROM golang:1.14-alpine as prod
RUN apk update && apk upgrade && \
  apk add --no-cache bash git openssh

RUN mkdir -p /app
WORKDIR /app
COPY . .

# Development
FROM prod as dev
RUN apk add --no-cache make
RUN go get -u github.com/cosmtrek/air
RUN mkdir -p /app
WORKDIR /app
COPY . .