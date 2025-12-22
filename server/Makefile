APP_NAME=server

dev:
	go run main.go

build:
	go build -o bin/$(APP_NAME)

test:
	go test ./...

lint:
	golangci-lint run

clean:
	rm -rf bin