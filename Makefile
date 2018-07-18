GOBUILD=go build
ENV=env GOOS=linux

all: clean build

build: main

vendor:
	govendor update github.com/pydio/pydio-sdk-go


main:
	go build -ldflags "-X github.com/pydio/pydio-sdk-go/config.version=${CELLS_VERSION} -X github.com/pydio/pydio-sdk-go/config.BuildStamp=`date -u +%Y-%m-%dT%H:%M:%S` -X github.com/pydio/pydio-sdk-go/config.BuildRevision=`git rev-parse HEAD`" -o pydio-sdk-go main.go

dev:
	go build -ldflags "-X github.com/pydio/pydio-sdk-go/config.version=0.2.0 -X github.com/pydio/pydio-sdk-go/config.BuildStamp=2018-01-01T00:00:00 -X github.com/pydio/pydio-sdk-go/config.BuildRevision=dev" -o pydio-sdk-go main.go

cleanall: stop clean rm

clean:
	rm -f pydio-sdk-go
