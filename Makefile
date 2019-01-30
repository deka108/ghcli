include .env
export

test-all: test-ghutil test-cmd-%

test-ghutil:
	@cd ghutil && go test -v

test-cmd-repo:
	@cd cmd && go test -v -run .*Repo.*

test-cmd-team:
	@cd cmd && go test -v -run .*Team.*

build:
	@go build -o bin/ghcli

run:
	./bin/ghcli

run-main: 
	@go run main.go

docker-image:
	@docker build -t deka108/ghcli .

export-env:
	@export $(cat .env | xargs) 