.DEFAULT_GOAL := help

BINARY=cvm
PROJECT=github.com/hawky-4s-/cvm

STATE=$(shell test -z "$$(git status --porcelain)" || echo '-dirty')
COMMIT=$(shell git rev-parse HEAD)
VERSION=$(shell git describe --exact-match --abbrev=0 2>/dev/null || echo 'dev')
GO_VERSION=$(shell go version | awk '{print $$3}' | cut -c3-)
LDFLAGS=-ldflags "-X $(PROJECT)/internal.commit=$(COMMIT)$(STATE) \
		-X $(PROJECT)/internal.version=$(VERSION) \
		-X $(PROJECT)/internal.goVersion=$(GO_VERSION)"

#SOURCES:=$(shell find . -name '*.go' -exec dirname {} \; | uniq | grep -v 'vendor')
SOURCES:=$(shell go list ./...)

################################ TARGETS ################################

.PHONY: all
all: fetch-dependencies check test build  ## run all targets

.PHONY: clean
clean: ## clean up build dir
	rm -rf bin
	go clean -x $(SOURCES)

.PHONY: dev
dev: clean tools
	realize start .

.PHONY: lint
lint: ## lint go files
	go fmt $(SOURCES)
	go fix $(SOURCES)

.PHONY: check
check: lint
	go vet -v $(SOURCES)

.PHONY: test
test: ## run tests
	@echo "" > coverage.txt
	@for d in $(SOURCES); do \
		go test -race -coverprofile=profile.out -covermode=atomic "$$d"; \
		if [ -f profile.out ]; then \
			cat profile.out >> coverage.txt; \
			rm profile.out; \
		fi \
	done

.PHONY: build
build: clean ## build and run binary
	mkdir -p bin
	go build $(LDFLAGS) -race -v -o bin/$(BINARY) ./cmd/$(BINARY)
	chmod u+x bin/$(BINARY)
	./bin/$(BINARY) version

.PHONY: install
install: ## install binary in GOPATH
	go install $(LDFLAGS) -v $(SOURCES)

.PHONY: tools
tools: ## install required tools
	@echo Current working dir: $$(pwd)
	@test $$(command -v dep) || curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	@test $$(command -v stringer) || go get -u -a golang.org/x/tools/cmd/stringer
	@cp .githooks/* .git/hooks/

.PHONY: fetch-dependencies
fetch-dependencies: tools ## fetch dependencies using dep
	dep ensure

################################ DOCKER ################################

.PHONY: docker-dev-req
docker-dev-req:
	docker build -t realize:1.11.1-alpine3.8 -f Dockerfile.realize .

.PHONY: docker-dev
docker-dev: clean
	docker run -it --rm --name $(BINARY) -v $$(pwd):/go/src/$(PROJECT) -p 5001:5001 realize:1.11.1-alpine3.8

.PHONY: docker-exec
docker-exec:
	docker exec -it $(BINARY) /bin/bash -c 'cd /go/src/$(PROJECT) && realize start .'

################################ MISC ################################

.PHONY: help
help: ## show help for targets
	grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
