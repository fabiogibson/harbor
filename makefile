SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

BINARY=harbor

VERSION=0.0.1
BUILD_TIME=`date +%FT%T%z`

LDFLAGS=-ldflags "-X github.com/fabiogibson/harbor/core.Version=${VERSION} -X github.com/fabiogibson/harbor/core.BuildTime=${BUILD_TIME}"

.DEFAULT_GOAL: $(BINARY)

$(BINARY): $(SOURCES)
	@go build -o ./bin/${BINARY} ./cli/main.go

.PHONY: install
install:
	@go install ${LDFLAGS} ./...

.PHONY: clean
clean:
	@rm -rf ./bin