include .env
export

build:
	go build -v ./cmd/gorestapi

start:
	./gorestapi

run:
	go run ./cmd/gorestapi

start-db:
	docker-compose up -d

migrate-status:
	sql-migrate status
migrate:
	sql-migrate up
migrate-down:
	sql-migrate down