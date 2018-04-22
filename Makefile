export GOPATH:=$(HOME)/.gopath:$(PWD)

test:
	@( clear && go vet src/*.go && go vet src/service/*.go && go vet src/data/*.go && cd test/unit && go test )
	@( golint src/... && golint test/... )

build: 
	@[ -d bin ] || mkdir bin
	go build -o bin/list-service src/main.go

install-deps:
	go get github.com/go-zoo/bone
	go get github.com/golang/lint/golint
	go get github.com/franela/goblin
	go get github.com/darrylwest/go-unique/unique
	go get -u github.com/darrylwest/cassava-logger/logger
	go get github.com/lib/pq

linux:
	@[ -d docker ] || mkdir docker
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o docker/list-service src/main.go

format:
	( gofmt -s -w src/*.go src/*/*.go test/*/*.go )

lint:
	@( golint src/... && golint test/... )

run:
	go run src/main.go --db-filename ~/database/list-service.db --port 9040

watch:
	go-watcher --loglevel=4

edit:
	make format
	vi -O3 src/*/*.go test/*/*.go src/*.go

docker-dev:
	docker run -it --name lister-dev --publish 9040:8080 --volume $(PWD):/opt ebay/debian-gcc:latest

.PHONY: format lint test qtest watch run test-hub test-worker docker-dev

