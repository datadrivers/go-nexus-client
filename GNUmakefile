GOCMD=go
GOTEST=$(GOCMD) test
GOBUILD=$(GOCMD) build

all: test build

test:
	$(GOTEST) -v -cover ./...

build:
	$(GOBUILD) -v ./...