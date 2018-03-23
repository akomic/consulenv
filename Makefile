GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=consulenv

.PHONY: list build build-linux clean

list:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | xargs
build:
	$(GOBUILD) -o $(BINARY_NAME)
build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -ldflags="-s -w" -o $(BINARY_NAME).linux.amd64
build-darwin:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -ldflags="-s -w" -o $(BINARY_NAME).darwin.amd64
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_NAME).linux.amd64
	rm -f $(BINARY_NAME).darwin.amd64

all: build build-linux clean
