golangci_version=v1.61.0
golangci_installed_version=$(shell golangci-lint version --format short 2>/dev/null)

lint-install:
ifneq ($(golangci_installed_version),$(golangci_version))
	@echo "--> Installing golangci-lint $(golangci_version)"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version)
endif

lint-fix:
	@echo "--> Running linter"
	$(MAKE) lint-install
	@./scripts/lint-all.sh --fix
