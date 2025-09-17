APP=api

.PHONY: up down run dev lint test migrate-up migrate-down seed swag

up:
	\tdocker compose up -d

down:
	\tdocker compose down -v

run:
	\tgo run ./cmd/api

dev:
	\tair || (go install github.com/air-verse/air@latest && air)

lint:
	\tgolangci-lint run || (go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest && golangci-lint run)

test:
	\tgo test ./... -count=1

migrate-up:
	\tmigrate -path infra/db/migrations -database "$(DB_DSN)" up

migrate-down:
	\tmigrate -path infra/db/migrations -database "$(DB_DSN)" down 1

seed:
	\tgo run ./scripts/seed

swagger:
	\techo "Keep OpenAPI in ./api/openapi.yaml (generate UI later)"