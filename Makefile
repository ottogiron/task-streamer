SOURCE_DIRS := $(shell go list ./... | grep -v /vendor/)

test: lint
	 @go test -v $(SOURCE_DIRS)


lint:
	@go fmt $(SOURCE_DIRS)
	@go vet $(SOURCE_DIRS)
