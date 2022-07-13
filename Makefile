PROJECTNAME=$(shell basename "$(PWD)")
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard cmd/search/main.go)
GIT=$(shell which git)
GOLANGCI_LINT=$(shell which golangci-lint run)
DOCKER_COMPOSE=$(shell which docker-compose)
export GO111MODULE=on
export GOPROXY=
export GOSUMDB=off


## build_release: Build linux native static linked release
build_release: go-clean go-mod-download go-test go-build

## lint: Start linter for project
lint:
	@echo " > Start golang linter"
	@$(GOLANGCI_LINT) run --timeout 5m0s

## run: Run compile version with example config
run:
	@GOOS=darwin GOARCH=arm64 go run $(GOFILES) config.example.yml

go-build: .build-main

.build-main:
	@echo "  >  Building binary version ${SERVICE_VERSION}..."
	@CGO_ENABLED=0 GOOS=linux GOBIN=$(GOBIN) go build -a -ldflags="-X 'main.Version=${SERVICE_VERSION}'" -installsuffix cgo -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)

go-generate:
	@echo "  >  Generating dependency files..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go generate ./...

go-mod-download:
	@echo "  >  Checking if there is any missing dependencies..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go mod download

## go-test: Testing current project code
go-test:
	@echo " > Testing ..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -v ./...

go-clean:
	@echo "  >  Cleaning build cache"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo


## protoc: Generate swagger & grpc interface files
.PHONY: protoc
protoc:
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
		--grpc-gateway_out=logtostderr=true:./api \
		--swagger_out=allow_merge=true,merge_file_name=api:./api \
		--go_out=plugins=grpc:./api ./api/*.proto