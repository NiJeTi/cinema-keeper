# cannot be migrated to go tool
# reason: https://golangci-lint.run/welcome/install/#install-from-sources
GOLANGCI_LINT_IMAGE=golangci/golangci-lint:v1.64-alpine

.PHONY: deps
deps:
	docker pull $(GOLANGCI_LINT_IMAGE)

.PHONY: mocks
mocks:
	rm -rf ./internal/generated/mocks
	go tool mockery

.PHONY: lint
lint:
	$(MAKE) deps

	docker run -t --rm -v $(PWD):/src -w /src $(GOLANGCI_LINT_IMAGE) golangci-lint run

.PHONY: test
test:
	./scripts/test.sh

.PHONY: debug
debug:
	docker compose \
		-f docker-compose.yaml -f docker-compose.override.yaml \
		-p cinema-keeper \
		up -d \
		db migrator

.PHONY: run
run:
	docker compose \
		-f docker-compose.yaml -f docker-compose.override.yaml \
		-p cinema-keeper \
		up -d \
		db migrator service

.PHONY: stop
stop:
	docker compose \
		-f docker-compose.yaml -f docker-compose.override.yaml \
		-p cinema-keeper \
		down -v --rmi local
