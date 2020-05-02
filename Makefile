# Local
install:
	@go get -u ./src
	@echo "All package installed"

run:
	@go run ./src/main.go

watch:
	@air -c air.conf

build:
	@go build -o ./build/main ./src

docker-dev:
	@docker-compose -f docker-compose.yml -f docker-compose.dev.yml up

docker-prod:
	@docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d

kill-port:
	@kill -9 $$(lsof -t -i:8080)
	@echo "Port 8080 is killed"

test:
	@go test ./__tests__/ -v -coverpkg=./... -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out
	@go tool cover -html=coverage.out
	@rm -f coverage.out

lint:
	@golangci-lint -E bodyclose,misspell,gocyclo,dupl,gofmt,golint,unconvert,goimports,depguard,gocritic,funlen,interfacer run