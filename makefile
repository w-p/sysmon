.PHONY: setup

VERSION=v0.0.1

clean:
	-rm ./bin/*
	-rm ./install/mac/Sysmon.app/Contents/MacOS/*
	-rm -fR ./vendor

setup:
	go mod tidy
	go mod vendor
	go install github.com/mgechev/revive@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest

lint:
	revive -config revive.toml -exclude vendor/... ./...
	go vet ./pkg
	staticcheck ./...

test: lint
	go test ./...

build: build-linux build-mac build-windows

build-linux:
	-mkdir ./bin
	GO111MODULE=on GOOS=linux GOARCH=amd64 go build \
		-o ./bin/sysmon-amd64-linux

build-mac:
	-mkdir ./bin
	GO111MODULE=on GOOS=darwin GOARCH=amd64 go build \
		-o ./bin/sysmon-amd64-darwin
	cp ./bin/sysmon-amd64-darwin ./install/mac/Sysmon.app/Contents/MacOS

build-windows:
	-mkdir ./bin
	GO111MODULE=on GOOS=windows GOARCH=amd64 go build \
		-ldflags -H=windowsgui \
		-o ./bin/sysmon-amd64-win.exe

release:
	gh release create $(VERSION) \
		-t "Sysmon $(VERSION)" \
		"./bin/sysmon-amd64-linux#Linux" \
		"./bin/sysmon-amd64-darwin#Mac OS" \
		"./bin/sysmon-amd64-win.exe#Windows"
