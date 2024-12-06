
APPNAME           := advent-of-code-2024-4
APP_GO_REPO_NAME  := github.com/fnickels/advent-of-code-2024-4


# check to see if awk is installed
AWK_INSTALLED := $(shell command -v awk 2> /dev/null)
ifeq ($(AWK_INSTALLED),)
$(error awk is not installed)
endif

# check to see if go is installed
GO_INSTALLED := $(shell command -v go 2> /dev/null)
ifeq ($(GO_INSTALLED),)
$(error go is not installed)
endif

# check to see if dlv is installed
DLV_INSTALLED := $(shell command -v dlv 2> /dev/null)
ifeq ($(DLV_INSTALLED),)
$(warning WARNING: Delve (dlv) is not installed, if you intend to use the go debugger, please install it with 'make delve_install')
endif

# check to see if golangci-lint is installed
GOLANGCI_LIST_INSTALLED := $(shell command -v golangci-lint 2> /dev/null)
ifeq ($(GOLANGCI_LIST_INSTALLED),)
$(warning WARNING: golangci-lint is not installed - please install, see README.md, or run 'make golangci_lint_install')
endif


##@ General

HELP_DISPLAY_COLUMS := 30

.PHONY: help
help: ## Shows this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-$(HELP_DISPLAY_COLUMS)s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) }' $(MAKEFILE_LIST)


## set default target if none is specified
## otherwise the First target in the file is the default target
.DEFAULT_GOAL := help


##@ Go

.PHONY: goinit
# goinit: export GOPATH=/mnt/c/git/
goinit: ## go init - Initialize Modules for this project
	-rm -rf go.mod go.sum
	go mod init $(APP_GO_REPO_NAME)
	go mod tidy


.PHONY: gotidy
gotidy:  ## go mod tidy - cleans up go modules
	go mod tidy

.PHONY: goversion
goversion: ## go & related tool versions 
	@go version
	@-dlv version
	@-golangci-lint --version

.PHONY: golint
golint:  ## Lint Go code
	-golangci-lint run


##@ Build  


.PHONY: run
run: goversion goinit golint gotest build exec ## Lint, Test, Build & Run the Go binary with test data



.PHONY: build
build: export GOARCH =
build: export GOOS   =
build: export CGO_ENABLED = 0
build: ## Build Go binary (current O/S)
	go build -o $(APPNAME) $(APP_GO_REPO_NAME)/src

.PHONY: debug
debug: export GOARCH =
debug: export GOOS   =
debug: export CGO_ENABLED = 0
debug: ## Build Go binary for Debugging (current O/S)
	go build -o $(APPNAME)-debug \
		-gcflags "all=-N -l" \
		$(APP_GO_REPO_NAME)/src


.PHONY: build-all
build-all: build-darwin build-linux build-windows ## Build Go binary for all OSes (Darwin, Linux, Windows)

.PHONY: gotest 
gotest: export GOARCH =
gotest: export GOOS   =
gotest: export CGO_ENABLED = 0
gotest: | ## Execute Go Unit & Integration tests 
	go test  ./...
	
##@ Execute Test Data

.PHONY: exec
exec: ## Run the Go binary with test data (current O/S)
	./$(APPNAME) 


####################
# Go Support Tools #
####################

##@ Go Support Tool Installers

.PHONY: golangci_lint_install
golangci_lint_install:  ## install latest version of golangci-lint
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: gofumptg_install
gofumptg_install:  ## install latest version of gofumpt
	go install mvdan.cc/gofumpt@latest

.PHONY: delve_install dlv_install
delve_install:  ## install latest released version of dlv
	@go install github.com/go-delve/delve/cmd/dlv@master

dlv_install: delve_install