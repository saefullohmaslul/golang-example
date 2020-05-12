# Production
FROM golang:1.14-alpine as prod
RUN apk update && apk upgrade && \
  apk add --no-cache bash git openssh

RUN mkdir -p /app
WORKDIR /app
COPY . .
EXPOSE 8080

# Development
FROM prod as local
RUN apk add --no-cache make
RUN go get -u github.com/cosmtrek/air
RUN mkdir -p /app
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 8080
CMD ["make", "watch"]