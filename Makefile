# Build Variables
BUILD_TIME 		?= $(shell date +%s)
BUILD_VERSION 	?= $(shell git rev-parse HEAD)

# Docker Configuration
DOCKER_IMAGE	?= gcr.io/soon-fm-production/trackstore
DOCKER_TAG		?= latest

# Go Compilation Flags
GOOUTDIR 		?= .
GOOS 			?=
GOARCH 			?=
CGO_ENABLED 	?= 0

# Bin Name
BIN_NAME 		?= trackstore
BIN_SUFFIX   	?=
ifneq ($(GOOS),)
ifneq ($(GOARCH),)
BIN_SUFFIX 		= .$(GOOS)-$(GOARCH)
endif
endif

# LDFlags
BUILD_TIME_LDFLAG 		?= -X trackstore/app.timestamp=$(BUILD_TIME)
BUILD_VERSION_LDFLAG 	?= -X trackstore/app.version=$(BUILD_VERSION)
LDFLAGS 				?= "$(BUILD_TIME_LDFLAG) $(BUILD_VERSION_LDFLAG)"

all: linux darwin

linux:
	docker run --rm -it -e GOOS=linux -e GOARCH=amd64 -v $(PWD):/go/src/trackstore $(DOCKER_IMAGE):onbuild build

darwin:
	docker run --rm -it -e GOOS=darwin -e GOARCH=amd64 -v $(PWD):/go/src/trackstore $(DOCKER_IMAGE):onbuild build

build:
	GOOS=$(GOOS) \
	GOARCH=$(GOARCH) \
	CGO_ENABLED=$(CGO_ENABLED) \
	go build \
		-v \
		-ldflags $(LDFLAGS) \
		-o "$(GOOUTDIR)/$(BIN_NAME)$(BIN_SUFFIX)"

test:
	go test -race -v -cover $(shell go list ./... | grep -v ./vendor/)

image:
	docker build --force-rm -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

onbuild:
	docker build --force-rm -f Dockerfile.onbuild -t $(DOCKER_IMAGE):onbuild .
