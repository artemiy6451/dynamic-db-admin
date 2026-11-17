APP_NAME=dynamic-ad-admin


.PHONY=run test build clean

run:
	go run .

test:
	go test ./...

build:
	go build -o build/$(APP_NAME) .

clean:
	rm -rf build
