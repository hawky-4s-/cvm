BINARY=cvm

LDFLAGS=-ldflags "-X github.com/hawky-4s-/cvm/internal.commit=`git rev-parse HEAD`"

SOURCEDIR=.
SOURCES:=$(shell find $(SOURCEDIR) -name '*.go')

.DEFAULT_GOAL: $(BINARY)

.PHONY: $(BINARY)
$(BINARY): clean
	mkdir -p bin
	go build ${LDFLAGS} -o bin/${BINARY} cmd/cvm/main.go
	chmod u+x bin/$(BINARY)

.PHONY: install
install:
	go install ${LDFLAGS} ./...

.PHONY: clean
clean:
	rm -rf bin

.PHONY: dev
dev: clean tools
	realize start .

.PHONY: docker-dev-req
docker-dev-req:
	docker build -t realize:1.11.1-alpine3.8 -f Dockerfile.realize .

.PHONY: docker-dev
docker-dev: clean
	docker run -it --rm -v $$(pwd):/app -p 5001:5001 realize:1.11.1-alpine3.8

.PHONY: tools
tools:
	@test $$(command -v dep) || curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	@test $$(command -v stringer) || go get -u -a golang.org/x/tools/cmd/stringer

.PHONY: release
release: require-GITHUB_TOKEN
	GOVERSION=$$(go version) goreleaser

.PHONY: release-draft
release-draft: require-GITHUB_TOKEN
	GOVERSION=$$(go version) goreleaser --snapshot

.PHONY: release-docker
release-docker:
	docker pull goreleaser/goreleaser release
	# https://github.com/goreleaser/goreleaser/blob/master/Dockerfile
	docker run --rm --privileged \
      -v $PWD:/go/src/github.com/user/repo \
      -v /var/run/docker.sock:/var/run/docker.sock \
      -w /go/src/github.com/user/repo \
      -e GITHUB_TOKEN \
      -e DOCKER_USERNAME \
      -e DOCKER_PASSWORD \
      goreleaser/goreleaser release

.PHONY: fetch-dependencies
fetch-dependencies: tools
	dep ensure

################################ MISC ################################
#.PHONY: require-%
#require-%:
#    @: $(if ${${*}},,$(error You must pass the $* environment variable))
#    @echo 'Had the variable (in make).'
