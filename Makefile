# Tải biến môi trường từ .eninclude .env

ifneq (,$(wildcard .env))
    include .env
    export
endif

ENV ?= dev

ifeq ($(ENV), dev)
	DB_HOST := $(DB_HOST_DEV)
else ifeq ($(ENV), prod)
	DB_HOST := $(DB_HOST_PROD)
endif

init: 
	@cd backend && go get github.com/olekukonko/tablewriter
	@cd backend && go get github.com/spf13/cobra 

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
	docker-compose --profile down
	docker-compose --profile dev up -d --build

# Chạy production
