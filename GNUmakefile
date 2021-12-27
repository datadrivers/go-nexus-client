GOCMD=go
GOTEST=$(GOCMD) test
GOBUILD=$(GOCMD) build

NEXUS_HOST=$(shell cd ./scripts && ./detect-docker-env-ip.sh)
MINIO_HOST=$(shell if [ "$(NEXUS_HOST)" = "127.0.0.1" ]; then echo "minio"; else echo "$(NEXUS_HOST)"; fi;)
NEXUS_PORT=$(shell grep -E "(NEXUS_PORT=)" ./scripts/.env | grep -oE "[0-9]+")
MINIKUBE_MOUNT_PID=$(word 1,$(shell ps | grep -v grep | grep 'minikube mount' | grep $(PWD)/scripts))

start-services:
ifeq (minikube,$(MINIKUBE_ACTIVE_DOCKERD))
ifeq (,$(MINIKUBE_MOUNT_PID))
	minikube mount $(PWD)/scripts:$(PWD)/scripts --uid=200 --gid=200 &
endif
endif
	cd ./scripts && ./start-services.sh && cd -

stop-services:
ifneq (,$(MINIKUBE_MOUNT_PID))
	kill $(MINIKUBE_MOUNT_PID)
endif
	cd ./scripts && ./stop-services.sh && cd -

restart-services: stop-services start-services

test:
	NEXUS_URL="http://$(NEXUS_HOST):$(NEXUS_PORT)" \
	NEXUS_USERNAME="admin" \
	NEXUS_PASSWORD="admin123" \
	AWS_ACCESS_KEY_ID="minioadmin" \
	AWS_SECRET_ACCESS_KEY="minioadmin" \
	AWS_ENDPOINT="http://$(MINIO_HOST):9000" \
	$(GOTEST) -v -cover ./...

build:
	$(GOBUILD) -v ./...

all: test build
