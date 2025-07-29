.PHONY: build run migrate test tidy lint clean docker-build docker-run docker-push ci

BINARY = student-api
VERSION ?= 1.0.0

include .env
export $(shell sed 's/=.*//' .env)

build:
	go build -o $(BINARY) .

run:
	go run .

migrate:
	PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -U $(DB_USER) -d $(DB_NAME) -f schema.sql

test:
	go test -v ./...

tidy:
	go mod tidy

lint:
	golangci-lint run

clean:
	rm -f $(BINARY)

docker-build:
<TAB>docker build -t nsahil992/student-api:$(VERSION) .

docker-run:
	docker run --env-file .env -p 8008:8080 nsahil992/student-api:$(VERSION)

docker-push:
	docker push nsahil992/student-api:$(VERSION)

ci: build test lint docker-build docker-push
