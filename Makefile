 DIR := ${CURDIR}
.DEFAULT_GOAL := help
.PHONY: test lint

test: ## Test the project
	cd test && go test -v

lint: ## Lint the project
	go fmt

build: ## Build the project
	go build

install: ## Build & Install as a binary to the PATH
	go install

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sed 's/Makefile://' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
