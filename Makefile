GOCMD ?= go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean -cache
GOGET = $(GOCMD) get
BINARY_NAME = ssh-automation


default: get build


get:
	@go mod tidy

build:
	CGO_ENABLED=0 $(GOBUILD) -o $(BINARY_NAME) -v -ldflags="-s -w"
	@upx $(BINARY_NAME)


test: build
	cp $(BINARY_NAME) test/$(BINARY_NAME)
	#cp config.yaml test/config.yaml


clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)


