# Local
install:
	@go get -u ./
	@echo "All package installed"

run::
	@go run main.go

kill-port:
	@kill -9 $$(lsof -t -i:8080)
	@echo "Port 8080 is killed"


# Docker utility
PID_FILE = /tmp/app.pid
GO_FILES = $(wildcard *.go)

start:
	go run $(GO_FILES) & echo $$! > $(PID_FILE)

stop:
	-kill `pstree -p \`cat $(PID_FILE)\` | tr "\n" " " |sed "s/[^0-9]/ /g" |sed "s/\s\s*/ /g"` 

before:
	@echo "STOPED app" && printf '%*s\n' "45" '' | tr ' ' _

restart: stop before start
	@echo "STARTED app" && printf '%*s\n' "45" '' | tr ' ' _

serve: start
	fswatch -or --event=Updated ./ | \
	xargs -n1 -I {} make restart

.PHONY: start before stop restart serve