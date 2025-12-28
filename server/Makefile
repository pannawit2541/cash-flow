APP_NAME=server
ENV_FILE=.env

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

migrate-up:
	@export $$(grep -v '^#' $(ENV_FILE) | xargs) && \
	migrate -database "$$MONGO_URI" -path migrations up

migrate-down:
	@export $$(grep -v '^#' $(ENV_FILE) | xargs) && \
	migrate -database "$$MONGO_URI" -path migrations down