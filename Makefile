.PHONY: build clean tool lint help

all: build

build:
	go build -v .

tool:
	go vet . | grep -v vendor; true
	gofmt -w .

lint:
	golint ./... | grep -v vendor; true

clean:
	rm -f go_webapp
	go clean -i .

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"
