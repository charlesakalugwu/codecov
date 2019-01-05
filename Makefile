COMMIT=$(shell git rev-parse --short HEAD)$(shell [[ $$(git status --porcelain --ignored) = "" ]] && echo -clean || echo -dirty)

# all is the default target to build everything
all: clean build

build: generate
	go build ./...

clean:
	rm -f coverage.out

test: unit

generate:
	go generate ./...

verify:
	go vet ./...

unit: generate
	go test ./... -coverprofile=coverage.out -covermode=atomic

cover: unit
	go tool cover -html=coverage.out

codecov: unit
	./hack/codecov-report.sh

.PHONY: clean verify unit
