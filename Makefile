.DEFAULT_GOAL := help

# ANSI color codes for better output
GREEN  := \033[32m
YELLOW := \033[33m
RED    := \033[31m
RESET  := \033[0m

# Global variables
GO_FILES_FMT := $(shell find . -type f -name '*.go' -not -path "./vendor/*" -not -path "*.gen.go" | tr "\n" " ")
GO_MAIN := "./cmd/"
BIN_SOURCE := "./bin/main"
BPATH := "./internal/basic-module/..."

.PHONY: build
build:
	@echo "$(YELLOW) Run build operation...$(RESET)"
	go build -o $(BIN_SOURCE) $(GO_MAIN)
	@echo "$(GREEN) Operation execution successfully.$(RESET)"

.PHONY: run
run: build
	@echo "$(YELLOW) Start app...$(RESET)"
	$(BIN_SOURCE) $(arg)
	@echo "$(GREEN) Operation execution successfully.$(RESET)"

.PHONY: test
test:
	@echo "$(YELLOW) Start test$(RESET)"
	go test -race -v -failfast -short -coverprofile=bin/coverage.out ./...
	@echo "$(YELLOW) Total coverage:$(RESET)"
	go tool cover -func=bin/coverage.out | grep total | awk '{print $$3}'
	@echo "$(GREEN) Operation execution successfully.$(RESET)"

.PHONY: btest
btest:
	@echo "$(YELLOW) Start bench test$(RESET)"
	go test -bench=. $(BPATH)
	@echo "$(GREEN) Operation execution successfully.$(RESET)"

.PHONY: bench
bench:
	@echo "$(YELLOW) Start bench sh test$(RESET)"
	./run_benchmarks.sh
	@echo "$(GREEN) Operation execution successfully.$(RESET)"

.PHONY: test-arg
test-arg:
	@echo "$(YELLOW)Running integration tests...$(RESET)"
	go test -tags $(hw) -count 1 -race -v ./...
	@echo "$(YELLOW) Total coverage:$(RESET)"
	go tool cover -func=bin/coverage.out | grep total | awk '{print $$3}'
	@echo "$(GREEN) Operation execution successfully.$(RESET)"

.PHONY: gen
gen:
	@echo "$(YELLOW) Start generation go files$(RESET)"
	go generate -v ./...
	@echo "$(GREEN) Operation execution successfully.$(RESET)"

# gofumpt -l -w . (https://github.com/mvdan/gofumpt)
.PHONY: fmt
fmt:
	@echo "$(YELLOW) Run gofumpt operation...$(RESET)"
	gofumpt -w $(GO_FILES_FMT)
	@echo "$(GREEN) Operation execution successfully.$(RESET)"

# GO MOD
.PHONY: vendor
vnd:
	@echo "$(YELLOW) Run vendor operation...$(RESET)"
	go mod vendor
	@echo "$(GREEN) Operation execution successfully.$(RESET)"

.PHONY: tidy
tidy:
	@echo "$(YELLOW) Run tidy operation...$(RESET)"
	go mod tidy
	@echo "$(GREEN) Operation execution successfully.$(RESET)"

.PHONY: vtidy
vtidy: tidy vnd

# GOLANGCI-LINT (https://golangci-lint.run/)
.PHONY: lint
lint:
	@echo "$(YELLOW) Run golangci-lint run operation...$(RESET)"
	golangci-lint run ./...
	@echo "$(GREEN) Operation execution successfully.$(RESET)"

.PHONY: lintfix
lintfix:
	@echo "$(YELLOW) Run golangci-lint run fix operation...$(RESET)"
	golangci-lint run ./... --fix
	@echo "$(GREEN) Operation execution successfully.$(RESET)"

.PHONY: lintf
lintf:
	@echo "$(YELLOW) Run golangci-lint run $(file) operation...$(RESET)"
	golangci-lint run $(file)
	@echo "$(GREEN) Operation execution successfully.$(RESET)"

.PHONY: lfmt
lfmt:
	@echo "$(YELLOW) Run golangci-lint fmt operation...$(RESET)"
	golangci-lint fmt
	@echo "$(GREEN) Operation execution successfully.$(RESET)"

.PHONY: help
help: 
	@echo "$(YELLOW)Available commands:$(RESET)"
	@echo "\t$(GREEN)fmt$(RESET):\t\toperation for start gofumpt command"
	@echo "\t$(GREEN)vnd$(RESET):\t\toperation for start vendor command"
	@echo "\t$(GREEN)tidy$(RESET):\t\toperation for start tidy command"
	@echo "\t$(GREEN)vtidy$(RESET):\t\toperation for start vendor+tidy command"
	@echo "\t$(GREEN)lint$(RESET):\t\toperation for start golangci-lint command"
	@echo "\t$(GREEN)lintfix$(RESET):\toperation for start golangci-lint fix command"
	@echo "\t$(GREEN)lint$(RESET):\t\toperation for start golangci-lint command: file = name of file or dir"
	@echo "\t$(GREEN)lfmt$(RESET):\t\toperation for start golangci-lint fmt command"
	@echo "\t$(GREEN)build$(RESET):\t\toperation for start go build command"
	@echo "\t$(GREEN)run$(RESET):\t\toperation for start $(BIN_SOURCE) command(arg=test)"
	@echo "\t$(GREEN)gen$(RESET):\t\toperation for start generation golang files command"
	@echo "\t$(GREEN)test$(RESET):\t\toperation for start test project command"
	@echo "\t$(GREEN)btest$(RESET):\t\toperation for start bench test project command(BPATH=./path/...)"
	@echo "\t$(GREEN)bench$(RESET):\t\toperation for start bench sh"
	@echo "\t$(GREEN)test-arg$(RESET):\toperation for start test project command (hw=hw01)"
