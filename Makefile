APP_NAME=dynamic-ad-admin


.PHONY: dev build package test test-race fmt vet check clean

dev:
	wails3 dev

build:
	wails3 build

package:
	wails3 package

test:
	go test ./...

test-race:
	go test -race ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

check: fmt vet test

clean:
	go clean
