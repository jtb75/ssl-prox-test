# Makefile for building Go project targeting Linux x86

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=myapp
BINARY_UNIX=$(BINARY_NAME)_unix

# Build directives
all: build
build:
	GOOS=linux GOARCH=386 $(GOBUILD) -o $(BINARY_NAME) -v
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	./$(BINARY_NAME)

# Dependency updates
deps:
	$(GOGET) -u
