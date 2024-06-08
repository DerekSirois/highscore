include .env

up:
	docker-compose down
	docker-compose up --build -d
	docker image prune -f

down:
	docker-compose down

migration-create:
	~/go/bin/goose -dir ${GOOSE_MIGRATION_DIR} ${GOOSE_DRIVER} ${GOOSE_DBSTRING} create $(name) sql

migration-up:
	~/go/bin/goose -dir ${GOOSE_MIGRATION_DIR} ${GOOSE_DRIVER} ${GOOSE_DBSTRING} up

test:
	@go test -v ./internal/...

run:
	@go run ./cmd/api/main.go