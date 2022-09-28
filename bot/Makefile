BINARY_NAME=iSamBot

all: clean build test

.DEFAULT_GOAL := help

.PHONY: fastrun
fastrun: ## fastrun: execute from main.go without compile
	@echo "> "$(BOT_TOKEN)
	go run cmd/isam/main.go

.PHONY: test
test: ## test: exute system test
	go test -v main.go

.PHONY: build
build: ## build: make binary version to test
	@echo "> building bot"
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=window go build -o ${BINARY_NAME}-windows main.go

.PHONY: run
run: ## run: execute binary version
	./${BINARY_NAME}-linux -token=$(BOT_TOKEN)

.PHONY: clean
clean: ## clean: execute `go clean` and remove ${BINARY_NAME}
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows

.PHONY: lint
lint: ## lint: execute lint
	golangci-lint run --enable-all


.PHONY: help
help : ## Makefile
	@printf "%s\n" "Useful targets:"
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  make %-15s\033[0m %s\n", $$1, $$2}'

# Variable outputting/exporting rules
var-%: ; @echo $($*)
varexport-%: ; @echo $*=$($*)
