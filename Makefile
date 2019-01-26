include .env
export

test-ghutil:
	@cd ghutil && go test -v

test-cmd:
	@cd cmd && go test -v

build:
	@go build -o bin/ghcli

run:
	./bin/ghcli

run-main: 
	@go run main.go