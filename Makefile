# Local
install:
	@go get -u ./src
	@go mod download
	@echo "All package installed"

run:
	@go run ./src/main.go

watch:
	@air -c air.conf

build:
	@go build -o ./build/main ./src

run-local:
	@docker-compose -f docker-compose.yml up --build

down-local:
	@docker-compose -f docker-compose.yml down

run-production:
	@docker-compose -f docker-compose.production.yml up --build -d

down-production:
	@docker-compose -f docker-compose.production.yml down

kill-port:
	@kill -9 $$(lsof -t -i:8080)
	@echo "Port 8080 is killed"

test:
	@sh scripts/test.sh

test-unit:
	@mkdir -p coverage
	@go test ./... -coverprofile=coverage/unit_test.txt

test-int:
	@mkdir -p coverage
	@go test ./tests -coverpkg=./... -coverprofile=coverage/integration_test.txt

lint:
	@golangci-lint -E bodyclose,misspell,gocyclo,dupl,gofmt,golint,unconvert,goimports,depguard,gocritic,funlen,interfacer run

seed:
	@go run src/helpers/seeder/seeder.go

migrate:
	@go run src/helpers/migrate/migration.go