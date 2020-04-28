# Local
install:
	@go get -u ./
	@echo "All package installed"

run:
	@go run main.go

watch:
	@air -c air.conf

kill-port:
	@kill -9 $$(lsof -t -i:8080)
	@echo "Port 8080 is killed"

test:
	@go test ./__tests__/ -v