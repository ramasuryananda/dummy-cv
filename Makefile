include .env

build:
	@go build -v -o dummy-cv ./cmd/*.go

run:
	@echo "RUN dummy-cv..."
	make build
	@./dummy-cv

test:
	@go test ./internal/...

test-clean:
	@go clean -testcache
	@go test ./internal/...

test-coverage:
	@go test ./internal/... -cover -v

test-coverage-html:
	@go test ./internal/... -coverprofile=coverage.out
	@go tool cover -html=coverage.out

migrate:
	@migrate -database ${DB_MIGRATION_CONNECTION} -path database/migrations up

migrate-rollback:
	@migrate -database ${DB_MIGRATION_CONNECTION} -path database/migrations down

seed:
	@migrate -database ${DB_SEEDER_CONNECTION} -path database/seeders up

seed-rollback:
	@migrate -database ${DB_SEEDER_CONNECTION} -path database/seeders down
