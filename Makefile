# Tải biến môi trường từ .env
include .env
export $(shell sed 's/=.*//' .env)

ENV ?= dev

ifeq ($(ENV), dev)
	DB_HOST := $(DB_HOST_DEV)
else ifeq ($(ENV), prod)
	DB_HOST := $(DB_HOST_PROD)
endif

# Lệnh để tạo migration mới từ schema Ent
migrate-create:
	@echo "Creating new migration..."
	@cd backend && go run github.com/golang-migrate/migrate/v4/cmd/migrate create -ext sql -dir migrations $(name)

# Lệnh chạy migration lên (up)
migrate-up:
	@echo "Running migrations up..."
	@cd backend && migrate -path migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

# Lệnh rollback migration (down)
migrate-down:
	@echo "Rolling back last migration..."
	@cd backend && migrate -path migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" down

# Lệnh kiểm tra trạng thái của migration
migrate-status:
	@echo "Checking migration status..."
	@cd backend && migrate -source file://migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" version

migrate-force:
	@echo "Forcing migration version to $(version)..."
	@cd backend && migrate -path migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" force $(version)


# Tạo code Ent từ schema
ent-gen:
	@cd backend && go run entgo.io/ent/cmd/ent generate ./ent/schema

# Chuẩn bị schema và migration
prepare:
	make ent-gen
	make migrate-up

run-backend-local:
	@cd backend && go run main.go
run-frontend-local:
	@cd frontend && npm run build && npx serve out 

# Chạy dev bằng Docker Compose
run-dev:
	ENV=dev DB_HOST=$(DB_HOST_DEV) docker-compose --profile dev up --build

# Chạy production
run-prod:
	ENV=prod DB_HOST=$(DB_HOST_PROD) docker-compose --profile prod up --build

run-dev-linux:
	ENV=dev DB_HOST=$(DB_HOST_DEV) docker compose --profile dev up --build

.PHONY: migrate migrate-up migrate-create migrate-status migrate-down ent-gen prepare run-dev run-prod
