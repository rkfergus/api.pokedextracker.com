DIRS     ?= $(shell find . -name '*.go' | grep --invert-match 'vendor' | xargs -n 1 dirname | sort --unique)
PKG_NAME ?= app

BFLAGS ?=
LFLAGS ?=
TFLAGS ?=

COVERAGE_PROFILE ?= coverage.out

default: build

.PHONY: setup
setup:
	@echo "--> Setting up"
	go get -u -v github.com/alecthomas/gometalinter github.com/golang/dep/cmd/dep
	gometalinter --install

.PHONY: install
install:
	@echo "---> Installing dependencies"
	dep ensure

.PHONY: build
build: install
	@echo "---> Building"
	go build -o ./bin/$(PKG_NAME) $(BFLAGS)

.PHONY: lint
lint:
	@echo "---> Linting... this might take a minute"
	gometalinter --vendor --tests --deadline=2m $(LFLAGS) $(DIRS)

.PHONY: test
test:
	@echo "---> Testing"
	GIN_MODE=release go test ./... -coverprofile $(COVERAGE_PROFILE) $(TFLAGS)

.PHONY: enforce
enforce:
	@echo "---> Enforcing coverage"
	./scripts/coverage.sh $(COVERAGE_PROFILE)

.PHONY: html
html:
	@echo "---> Generating HTML coverage report"
	go tool cover -html $(COVERAGE_PROFILE)

.PHONY: clean
clean:
	@echo "---> Cleaning"
	rm -rf ./bin ./vendor
