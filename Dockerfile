# development
FROM golang:1.14-alpine AS dev

RUN apk add --no-cache make git gcc libc-dev
RUN go get -u github.com/cosmtrek/air
RUN go get -v github.com/rubenv/sql-migrate/...
RUN mkdir -p /app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 1323

CMD ["make", "dev"]

# builder image
FROM golang:1.14-alpine AS builder

RUN apk add --no-cache make
RUN mkdir -p /app

WORKDIR /app

COPY --from=dev /app ./

RUN make build

# production
FROM golang:1.14-alpine AS prod

RUN apk add --no-cache make
RUN mkdir -p /app

WORKDIR /app

COPY --from=builder /app/dist ./dist
COPY --from=builder /app/db ./db
COPY --from=builder /app/Makefile ./Makefile

EXPOSE 1323

CMD ["make", "start"]