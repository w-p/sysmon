.PHONY: setup build run

setup:
	go mod tidy
	go mod vendor

build:
	env GO111MODULE=on go build -o ./bin

run:
	./bin/sysmon

