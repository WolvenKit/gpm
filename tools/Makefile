 DIR := ${CURDIR}
.DEFAULT_GOAL := help

lint: ## Lint the go project
	go fmt

build: lint ## Build the go project
	echo "==> Building..."
	go build &&  \
	go install

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sed 's/Makefile://' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
