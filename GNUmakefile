GOCMD=go
GOTEST=$(GOCMD) test
GOBUILD=$(GOCMD) build

start-services:
	cd ./scripts && ./start-services.sh && cd -

stop-services:
	cd ./scripts && ./stop-services.sh && cd -

test:
	$(GOTEST) -v -cover ./...

build:
	$(GOBUILD) -v ./...

all: test build
