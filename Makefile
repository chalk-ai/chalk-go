.PHONY: all test

SHELL := /bin/bash
TEMP_FILE := $(shell mktemp)

help:  ## Show this help.
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'

test:  ## Run tests. Needs bash > 4.0, gotestsum, panicparse, and script
	@which gotestsum > /dev/null || go get gotest.tools/gotestsum@latest
	@which panicparse > /dev/null || go install github.com/maruel/panicparse/v2@latest
	@bash -c "CGO_ENABLED=0 GOTRACEBACK=all script -q /dev/null gotestsum --hide-summary=skipped --format-hide-empty-pkg -- -vet=all -shuffle=on ./... |& panicparse -rel-path"

release:  ## Release the go client
	@bash ./release.sh

cloc:
	cloc . --vcs=git --exclude-lang JSON,SVG,.pyi --not-match-f generated.go

.DEFAULT:
	@echo Unknown command. Available commands below
	@echo
	@make help
