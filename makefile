.DEFAULT_GOAL := seerlink-build

ENVIRONMENT ?= release

BUILDPATH ?= ./build/bin
BUILDER ?= SeerLink/builder
REPO := SeerLink/seerlink
COMMIT_SHA ?= $(shell git rev-parse HEAD)
VERSION = $(shell cat VERSION)
GO_LDFLAGS := $(shell tools/bin/ldflags)
GOFLAGS = -ldflags "$(GO_LDFLAGS)"

.PHONY: seerlink-build
seerlink-build:
	go build $(GOFLAGS) -o $(BUILDPATH)/seerlink ./core/

# Format for CI
.PHONY: presubmit
presubmit:
	goimports -w ./core
	gofmt -w ./core
	go mod tidy


