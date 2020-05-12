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

# test:
# 	@go test ./__tests__/ -v -coverpkg=./... -coverprofile=coverage.out.tmp ./...
# 	@cat coverage.out.tmp | grep -v "app/application.go" | grep -v "database/" | grep -v "package/" > coverage.out
# 	@go tool cover -func=coverage.out
# 	@go tool cover -html=coverage.out
# 	@rm -f coverage.out coverage.out.tmp

test:
	@make test-unit
	@make test-int
	@grep 'mode' coverage_unit.out > coverage.out.tmp; grep 'github' coverage_unit.out >> coverage.out.tmp; grep 'github' coverage_int.out >> coverage.out.tmp
	@cat coverage.out.tmp | grep -v "app/application.go" | grep -v "database/" > coverage.out
	@go tool cover -func=coverage.out
	@go tool cover -html=coverage.out
	@rm -f coverage*

test-unit:
	@go test ./... -coverprofile=coverage_unit.out

test-int:
	@go test ./tests -coverpkg=./... -coverprofile=coverage_int.out

lint:
	@golangci-lint -E bodyclose,misspell,gocyclo,dupl,gofmt,golint,unconvert,goimports,depguard,gocritic,funlen,interfacer run