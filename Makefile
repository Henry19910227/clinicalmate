.PHONY: run build test lint mysql-up mysql-down mysql-reset

run:
	go run cmd/main.go -config config.yaml

build:
	go build -o bin/clinicalmate cmd/main.go

test:
	go test ./...

lint:
	golangci-lint run

mysql-up:
	docker compose up -d mysql

mysql-down:
	docker compose down

mysql-reset:
	docker compose down -v && docker compose up -d mysql
