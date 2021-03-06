#一定不能用4个空格代替tab

GOCMD=go
GOBUILD=$(GOCMD) build  -i -v -ldflags '-w -s'
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test ./...
GOGET=$(GOCMD) get -u -v

OS := $(shell uname -s | awk '{print tolower($$0)}')

BINARY = bin/go-server-cli

GOARCH = amd64

LDFLAGS = -ldflags="$$(govvv -flags)"

all: clean build

test:
	$(GOTEST)

lint:
	golint

build:
	env CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(GOARCH) GIN_MODE=release $(GOBUILD)  -o $(BINARY)

clean:
	$(GOCLEAN)
	@rm -f $(BINARY)