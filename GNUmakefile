GOCMD=go
GOTEST=$(GOCMD) test
GOBUILD=$(GOCMD) build

start-services:
	cd ./scripts && ./start-services.sh && cd -

stop-services:
	cd ./scripts && ./stop-services.sh && cd -

restart-services: stop-services start-services

test:
	$(GOTEST) -v -cover ./...

build:
	$(GOBUILD) -v ./...

all: test build
