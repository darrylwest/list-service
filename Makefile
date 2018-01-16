export GOPATH:=$(HOME)/.gopath:$(PWD)

test:
	@( go vet src/*.go )
	@( go vet src/*/*.go && cd test/unit && go test )

build: 
	@[ -d bin ] || mkdir bin
	go build -o bin/list-service src/main.go

install-deps:
	go get github.com/golang/lint/golint
	go get github.com/franela/goblin
	go get github.com/darrylwest/go-unique/unique
	go get github.com/boltdb/bolt
	go get -u github.com/darrylwest/cassava-logger/logger

linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o docker/list-service src/main.go &&

format:
	( gofmt -s -w src/*.go src/*/*.go test/*/*.go )

lint:
	@( golint src/... && golint test/... )

run:
	go run src/main.go

watch:
	go-watcher --loglevel=4

edit:
	make format
	vi -O3 src/*/*.go test/*/*.go src/*.go

.PHONY: format lint test qtest watch run test-hub test-worker

