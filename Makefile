include .env
SHELL := /bin/bash

LOCAL_MIGRATION_DIR=./internal/db/migrations
LOCAL_MIGRATION_DSN="host=localhost port=$(PG_PORT) dbname=$(PG_DATABASE_NAME) user=$(PG_USER) password=$(PG_PASSWORD) sslmode=disable"

build:
	go build -o ./bin/keeper ./cmd/keeper

mig-s:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

mig-u:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

mig-d:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

gen:
	make gen-auth-api
	make gen-storage-api

gen-auth-api:
	mkdir -p internal/gen/auth/
	protoc --proto_path api/auth \
	--go_out=internal/gen/auth/ --go_opt=paths=source_relative \
	--go-grpc_out=internal/gen/auth/ --go-grpc_opt=paths=source_relative \
	api/auth/auth.proto

gen-storage-api:
	mkdir -p internal/gen/storage/
	protoc --proto_path api/storage \
	--go_out=internal/gen/storage/ --go_opt=paths=source_relative \
	--go-grpc_out=internal/gen/storage/ --go-grpc_opt=paths=source_relative \
	api/storage/storage.proto