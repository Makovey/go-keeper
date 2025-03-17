include .env
SHELL := /bin/bash

LOCAL_BIN:=$(CURDIR)/bin
LOCAL_MIGRATION_DIR=./internal/db/migrations
LOCAL_MIGRATION_DSN="host=localhost port=$(PG_PORT) dbname=$(PG_DATABASE_NAME) user=$(PG_USER) password=$(PG_PASSWORD) sslmode=disable"

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest
	GOBIN=$(LOCAL_BIN) go install github.com/golang/mock/mockgen@v1.6.0

mig-s:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

mig-u:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

mig-d:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

gen-auth-api:
	mkdir -p internal/gen/auth/
	protoc --proto_path api/auth \
	--go_out=internal/gen/auth/ --go_opt=paths=source_relative \
	--go-grpc_out=internal/gen/auth/ --go-grpc_opt=paths=source_relative \
	api/auth/auth.proto