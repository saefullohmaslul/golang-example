dev:
	air -c .air.toml

build:
	go build -o ./dist/main ./internal

start:
	chmod +x ./dist/main
	./dist/main

migrate:
	sql-migrate up

migrate-undo:
	sql-migrate down

test:
	mkdir -p coverage
	go test -v -coverprofile ./coverage/cover.out ./...
	go tool cover -html=./coverage/cover.out -o ./coverage/cover.html
	open ./coverage/cover.html