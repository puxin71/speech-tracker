PROJECTNAME=$(shell basename "$(PWD)")

# Go related variables.
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard *.go)
RESOURCES="$(PWD)/resources"

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

compile: clean get build

build:
	@echo "  >  Building binary..."
	@go build -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)

generate:
	@echo "  >  Generating dependency files..."
	@go generate $(generate)

get:
	@echo "  >  Checking if there is any missing dependencies..."
	@go get $(get)

install:
	@go install $(GOFILES)

clean:
	@echo "  >  Cleaning build cache"
	@go clean

test:
	@go test -v ./...

start-server: 
	@RESOURCES=$(RESOURCES) $(GOBIN)/$(PROJECTNAME)
