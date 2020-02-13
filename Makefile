SOURCE_FILES := ./...
ROOT := .

export GO111MODULE := on
export GOBIN := $(shell pwd)/bin
export PATH := $(GOBIN):$(PATH)

.PHONY: setup
setup:
	go mod download
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.23.6
	curl -sSfL https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | sh -s v0.126.0

.PHONY: build
build:
	go build -o bin/barb cmd/barb/main.go

.PHONY: lint
lint:
	go fmt $(SOURCE_FILES)
	./bin/golangci-lint run --config .golangci.yml --fix $(SOURCE_FILES)

.PHONY: test
test:
	go test -race -coverprofile=coverage.txt -coverpkg=$(SOURCE_FILES) -covermode=atomic $(SOURCE_FILES) -run $(ROOT)

.PHONY: tag
tag:
	git tag $(VERSION)
	git push origin $(VERSION)

.PHONY: remove_tag
remove_tag:
	git tag -d $(VERSION)
	git fetch
	git push origin --delete $(VERSION)
	git tag -d $(VERSION)

.PHONY: docker_image
docker_image:
	docker build -t barb -f Dockerfile $(ROOT)

.PHONY: docker_publish
docker_publish:
	make docker_image
	docker tag barb:latest cathalmullan/barb:${TRAVIS_TAG}
	docker push cathalmullan/barb:${TRAVIS_TAG}
	docker tag barb:latest cathalmullan/barb:latest
	docker push cathalmullan/barb:latest

.PHONY: package
package:
	./bin/goreleaser

# TODO: Add GoDownloader - https://github.com/goreleaser/godownloader/pull/166
#.PHONY: install
#install:
#	./bin/godownloader -repo CathalMullan/barb > ./install.sh
