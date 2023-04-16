# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build #-race
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
OUT=tpm

.PHONY: all clean test run build install update-dependencies

all: build
clean:
	$(GOCLEAN) && rm -rf bin/$(OUT)
test:
	$(GOTEST) ./...
run:
	$(GORUN) main.go
build:
	$(GOBUILD) -o bin/$(OUT) main.go
install: build
	mkdir -p /usr/local/bin && cp bin/$(OUT) /usr/local/bin/$(OUT) && chmod +x /usr/local/bin/$(OUT)
update-dependencies:
	$(GOCMD) get -u all && $(GOCMD) mod tidy
