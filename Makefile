BIN_DIR = bin

run: ## Build project
	go run ./monnify.go

build: ## Build project
	go build -o ./bin/monnify ./monnify

test: ## Launch tests
	go test -v ./...

test-coverage: ## Launch tests
	go test -cover ./...

test-coverage-report: ## Launch tests
	go test -coverprofile=coverage.out ./... 

view-coverage-report: ## Launch tests
	go tool cover -html=coverage.out

bump: ## Update packages version
	go get -u ./...

about: ## Display info related to the build
	@echo "OS: ${OS}"
	@echo "Shell: ${SHELL} ${SHELL_VERSION}"
	@echo "Protoc version: $(shell protoc --version)"
	@echo "Go version: $(shell go version)"
	@echo "Go package: ${PACKAGE}"
	@echo "Openssl version: $(shell openssl version)"

help: ## Show this help
	@${HELP_CMD}
