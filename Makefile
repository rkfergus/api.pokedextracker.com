DIRS     ?= $(shell find . -name '*.go' | grep --invert-match 'vendor' | xargs -n 1 dirname | sort --unique)
PKG_NAME ?= app

BFLAGS ?=
LFLAGS ?=
TFLAGS ?=

COVERAGE_PROFILE ?= coverage.out

PSQL := $(shell command -v psql 2> /dev/null)

DATABASE_USER             ?= pokedex_tracker_user
TEST_DATABASE_NAME        ?= pokedex_tracker_test
DEVELOPMENT_DATABASE_NAME ?= pokedex_tracker

default: build

.PHONY: build
build: install
	@echo "---> Building"
	go build -o ./bin/$(PKG_NAME) $(BFLAGS)

.PHONY: clean
clean:
	@echo "---> Cleaning"
	rm -rf ./bin ./vendor

.PHONY: enforce
enforce:
	@echo "---> Enforcing coverage"
	./scripts/coverage.sh $(COVERAGE_PROFILE)

.PHONY: html
html:
	@echo "---> Generating HTML coverage report"
	go tool cover -html $(COVERAGE_PROFILE)

.PHONY: install
install:
	@echo "---> Installing dependencies"
	dep ensure

.PHONY: lint
lint:
	@echo "---> Linting"
	gometalinter --vendor --tests $(LFLAGS) $(DIRS)

.PHONY: migrate
migrate:
	@echo "---> Migrating"
	go run cmd/migrations/*.go migrate

.PHONY: rollback
rollback:
	@echo "---> Rolling back"
	go run cmd/migrations/*.go rollback

.PHONY: setup
setup:
	@echo "--> Setting up"
	go get -u -v github.com/alecthomas/gometalinter github.com/golang/dep/cmd/dep
	gometalinter --install
ifdef PSQL
	dropdb --if-exists $(DEVELOPMENT_DATABASE_NAME)
	dropdb --if-exists $(TEST_DATABASE_NAME)
	dropuser --if-exists $(DATABASE_USER)
	createuser --createdb $(DATABASE_USER)
	createdb -U $(DATABASE_USER) $(TEST_DATABASE_NAME)
	createdb -U $(DATABASE_USER) $(DEVELOPMENT_DATABASE_NAME)
else
	$(info Skipping database setup)
endif

.PHONY: test
test:
	@echo "---> Testing"
	GIN_MODE=release ENVIRONMENT=test go test ./... -coverprofile $(COVERAGE_PROFILE) $(TFLAGS)
