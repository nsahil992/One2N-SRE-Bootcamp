.PHONY: build run migrate test tidy lint

BINARY = student-api

include .env
export $(shell sed 's/=.*//' .env)

build:
	go build -o $(BINARY)

run:
	go run .

# migrate:
	#	psql -h $(DB_HOST) -U $(DB_USER) -d $(DB_NAME) -f schema.sql

test:
	go test -v ./...

tidy:
	go mod tidy

lint:
	golangci-lint run

clean:
	rm -f $(BINARY)