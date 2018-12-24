.SILENT:
export GOPATH:=$(HOME)/.gopath:$(PWD)

## help: show this help message
help:
	@( echo "" && echo "Makefile targets..." && echo "" )
	@( cat Makefile | grep '^##' | sed -e 's/##/ -/' | sort && echo "" )

## test: run project unit tests including lint
test:
	@( clear && go vet src/*.go && go vet src/app/*.go && cd test/unit && go test )
	@( golint src/... && golint test/... )

## qtest: run project unit tests without lint
qtest:
	@( clear && go vet src/*.go && go vet src/app/*.go && cd test/unit && go test )

## build: compile the project
build: 
	@[ -d bin ] || mkdir bin
	packr build -o bin/list-service src/main.go

## install-deps: download all project dependencies
install-deps:
	go get github.com/golang/lint/golint
	go get github.com/franela/goblin
	go get github.com/darrylwest/go-unique/unique
	go get github.com/darrylwest/cassava-logger/logger
	go get github.com/go-zoo/bone
	go get -u github.com/gobuffalo/packr/...
	go get github.com/lib/pq

## linux: compile to linux compatible with scratch or alpine docker
linux:
	@[ -d docker ] || mkdir docker
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o docker/list-service src/main.go

## format: format all the souce code
format:
	( gofmt -s -w src/*.go src/*/*.go test/*/*.go )

## lint: run lint on the project
lint:
	@( golint src/... && golint test/... )

## run: run the project at a specified port (9040)
run:
	packr build -o bin/list-service src/main.go && ./bin/list-service --port 9040

## status: curl script to fetch the status of a running instance
status:
	curl http://localhost:9040/api/list/status

## watch: run the standard watcher
watch:
	go-watcher --loglevel=4

## edit: edit the project (vi)
edit:
	make format
	vi -O3 src/*/*.go test/*/*.go src/*.go

## docker-dev: run the appliction in a container
docker-dev:
	docker run -it --name lister-dev --publish 9040:8080 --volume $(PWD):/opt ebay/debian-gcc:latest

.PHONY: format lint test qtest watch run test-hub test-worker docker-dev

