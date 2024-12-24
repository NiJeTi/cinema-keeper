MOCKERY_IMAGE=vektra/mockery:v2.50
GOLANGCI_LINT_IMAGE=golangci/golangci-lint:v1.62-alpine

.PHONY: deps
deps:
	docker pull $(MOCKERY_IMAGE)
	docker pull $(GOLANGCI_LINT_IMAGE)

.PHONY: mocks
mocks:
	$(MAKE) deps

	rm -rf ./internal/generated/mocks
	docker run -t --rm -v $(PWD):/src -w /src $(MOCKERY_IMAGE)

.PHONY: lint
lint:
	$(MAKE) deps

	docker run -t --rm -v $(PWD):/src -w /src $(GOLANGCI_LINT_IMAGE) golangci-lint run -v

.PHONY: devenv
devenv:
	docker compose -f docker-compose.yaml -f docker-compose.dev.yaml -p cinema-keeper up -d
