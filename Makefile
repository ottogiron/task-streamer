NAME := wizard
VERSION := 0.3.0
DATE := $(shell date)
COMMIT_ID := $(shell git rev-parse --short HEAD 2>/dev/null || echo $$COMMIT_ID)
GIT_REPO := $(shell git config --get remote.origin.url)
LD_FLAGS := '-X "git.corp.xoom.com/projects/inf/wizard/cmd.buildVersion=$(VERSION)" -X "git.corp.xoom.com/projects/inf/wizard/cmd.buildCommit=$(COMMIT_ID)" -X "git.corp.xoom.com/projects/inf/wizard/cmd.buildDate=$(DATE)"'
EXTRA_BUILD_VARS := CGO_ENABLED=0 GOARCH=amd64
SOURCE_DIRS := $(shell go list ./... | grep -v /vendor/)
all:  binaries

build-release: container

lint:
	@go fmt $(SOURCE_DIRS)
	@go vet $(SOURCE_DIRS)

test: lint
	 @go test -v $(SOURCE_DIRS)

binaries: test
	GOOS=darwin $(EXTRA_BUILD_VARS)  go build -ldflags $(LD_FLAGS) -o $(NAME)-darwin
	GOOS=linux $(EXTRA_BUILD_VARS)  go build -ldflags $(LD_FLAGS) -o $(NAME)-linux

image: binaries
	docker-flow build -f docker/Dockerfile

push: image
	docker-flow push

package-brew: binaries
	cp wizard-darwin wizard
	tar czf wizard.tgz wizard

package-linux: binaries
	cp wizard-linux wizard
	tar czf wizard.tgz wizard
