SHELL := /bin/bash

# -----------------
# Project Variables
# -----------------
PROJECT_NAME:=svc
PROJECT_ROOT?=$(shell pwd)
PROJECT_WORKDIR?=${PROJECT_ROOT}
PROJECT_CONFIG:=.env
PROJECT_DOCKERFILE_DIR?=${PROJECT_ROOT}/build/svc
OUTPUT_DIR:=${PROJECT_ROOT}/bin
BINARY_NAME:=svc
SCRIPTS_DIR := ${PROJECT_ROOT}/tools

# ---------------
# Command Aliases
# ---------------
GO_CMD:=go
GO_BUILD:=${GO_CMD} build
GO_MOD:=${GO_CMD} mod
GO_CLEAN:=${GO_CMD} clean
GO_GET:=${GO_CMD} get
GO_INSTALL:=${GO_CMD} install
DOCKER_CMD:=docker

# ---
# App Service
# ---
PROJECT_MAIN_PKG=cmd/${BINARY_NAME}
PROJECT_ENV_FILES:=$(addprefix ${PROJECT_ROOT}/,${PROJECT_CONFIG})

# -------------------
# Migration Variables
# -------------------
MIGRATION_DIR := ${PROJECT_ROOT}/migrations
MIGRATION_SRC_DIR := ${MIGRATION_DIR}/sql
MIGRATION_URL = "${DB_DRIVER}://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"
MIGRATION_CMD := migrate -source "file://${MIGRATION_SRC_DIR}" -database ${MIGRATION_URL}

# ----------------------
# Debug Build App
# ----------------------
DEBUG_DIR:=${OUTPUT_DIR}/debug
DEBUG_BIN:=${DEBUG_DIR}/${BINARY_NAME}
DEBUG_ENV_FILES:=$(addprefix ${DEBUG_DIR}/,${PROJECT_CONFIG})

# ----------------------
# Debug Build App
# ----------------------
COMPILE_DIR:=${OUTPUT_DIR}/compile
COMPILE_BIN:=${COMPILE_DIR}/${BINARY_NAME}

# -----------
# Additional function
# -----------
# Temp Variables
DOCKER_COMPOSE_LOCAL=docker-compose-local.yml
CMD_DOCKER_COMPOSE_LOCAL:=docker-compose -f ${DOCKER_COMPOSE_LOCAL} --env-file $(PROJECT_CONFIG)

# ------------
# Command List
# ------------

## help: Show command help
.PHONY: help
all: help
help: Makefile
	@echo
	@echo " List command you can use in "${PROJECT_NAME}":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## swagger-setup: Setup for swag command
.PHONY: swagger-setup
swagger-setup:
	@-echo "  > Setup swagger cli..."
	@-export GOBIN="$(go env GOPATH)/bin"
	@-${GO_GET} -u "github.com/swaggo/swag/cmd/swag"
	@-export PATH="$(go env GOPATH)/bin:$PATH"

## swagger: Generate swagger file base on annotation
.PHONY: swagger
swagger:
	@-echo "  > Generate swagger"
	@swag init -g ./cmd/${PROJECT_NAME}/main.go -o ./docs

## configure: Configure project file
.PHONY: configure
configure: --copy-env vendor
	@-echo "  > Configure: Done"
	@-export $(grep -v '^#' .env | xargs -d '\n')

.PHONY: --copy-env
--copy-env:
	@-echo "  > Copy .env (did not overwrite existing file)..."
	@-cp -n $(PROJECT_ROOT)/configs/.env.example $(PROJECT_CONFIG)

## migration: Make migration on migration sql
.PHONY: migration
migration:
	@read -p "Enter migration name:" migration; \
	${MIGRATION_CMD} create -ext sql -dir "${MIGRATION_SRC_DIR}" -seq $$migration;

## db-status: Prints the details and status information about all the migrations.
.PHONY: db-status
db-status:
	@${MIGRATION_CMD} version

## db-up: Run migration database upgrade
.PHONY: db-up
db-up:
	@-echo "  > Running up scripts..."
	@${MIGRATION_CMD} up

## db-down: Undo to previous migration version
.PHONY: db-down
db-down:
	@${MIGRATION_CMD} down 1

## db-clean: Clean database
.PHONY: db-clean
db-clean: --clean-prompt
	@-echo "  > Cleaning database..."
	@${MIGRATION_CMD} drop

.PHONY: --clean-prompt
--clean-prompt:
	@-echo -n "Are you sure want to clean all data in database? [y/N] " && read ans && [ $${ans:-N} = y ]

## vendor: Download dependencies to vendor folder
.PHONY: vendor
vendor: go.mod
	@-echo "  > Vendoring..."
	@${GO_MOD} vendor
	@-echo "  > Vendoring: Done"

## container-setup: Create containers dependencies for local development
.PHONY: container-setup
container-setup:
	@-echo " > Creating database..."
	@-$(CMD_DOCKER_COMPOSE_LOCAL) up -d db

## serve-docker: Create docs for local development
.PHONY: serve-docker
serve-docker:
	@-$(CMD_DOCKER_COMPOSE_LOCAL) build
	@-$(CMD_DOCKER_COMPOSE_LOCAL) up -d ${PROJECT_NAME}

## down-docker: Down docker at environment docker
.PHONY: down-docker
down-docker:
	@-$(CMD_DOCKER_COMPOSE_LOCAL) down

## re-run: Rebuild docker compose project
.PHONY: re-run
re-run:
	@-$(CMD_DOCKER_COMPOSE) up -d ${PROJECT_NAME} --build

## compile: Compile binary for deployment.
.PHONY: compile
compile: vendor
	@-echo "  > Compiling..."
	@CGO_ENABLED=0 GOOS=linux ${GO_BUILD} -gcflags="all=-N -l" -a -v \
		-o ${COMPILE_DIR}/${BINARY_NAME} ${PROJECT_ROOT}/${PROJECT_MAIN_PKG}
	@-echo "  > Copying required file..."
	@-echo "  > Output: ${COMPILE_DIR}"
	@-ls -la ${COMPILE_DIR}

.PHONY: --dev-build
--dev-build:
	@-echo "  > Compiling..."
	@${GO_BUILD} -o ${DEBUG_BIN} ${PROJECT_ROOT}/${PROJECT_MAIN_PKG}
	@-echo "  > Output: ${DEBUG_BIN}"

${DEBUG_ENV_FILES}: $(PROJECT_ENV_FILES)
	@-echo "  > Copying environment files..."
	@cp -R ${PROJECT_ENV_FILES} ${DEBUG_DIR}
