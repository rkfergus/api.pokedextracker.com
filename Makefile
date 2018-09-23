DIRS     ?= $(shell find . -name '*.go' | grep --invert-match 'vendor' | xargs -n 1 dirname | sort --unique)
PKG_NAME ?= app

BFLAGS ?=
LFLAGS ?=
TFLAGS ?=

COVERAGE_PROFILE ?= coverage.out

PSQL := $(shell command -v psql 2> /dev/null)

DATABASE_USER             ?= pokedex_tracker_admin
TEST_DATABASE_NAME        ?= pokedex_tracker_test
DEVELOPMENT_DATABASE_NAME ?= pokedex_tracker

default: build

.PHONY: build
build: install
	@echo "---> Building"
	CGO_ENABLED=0 go build -o ./bin/$(PKG_NAME) -installsuffix cgo -ldflags '-w -s' $(BFLAGS) ./cmd/serve

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

.PHONY: seed
seed:
	@echo "---> Populating seeds"
	go run cmd/seeds/*.go

.PHONY: serve
serve:
	@echo "---> Serving"
	gin --port 8648 --appPort 8649 --path . --build ./cmd/serve --immediate --bin ./bin/gin-$(PKG_NAME) run

.PHONY: setup
setup:
	@echo "--> Setting up"
	go get -u -v github.com/alecthomas/gometalinter github.com/golang/dep/cmd/dep github.com/codegangsta/gin
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
	GIN_MODE=release ENVIRONMENT=test go test -race ./pkg/... -coverprofile $(COVERAGE_PROFILE) $(TFLAGS)
