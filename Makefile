MOCKERY_VERSION=github.com/vektra/mockery/v2@v2.50
GOLANGCI_LINT_IMAGE=golangci/golangci-lint:v1.63-alpine
GOOSE_VERSION=github.com/pressly/goose/v3/cmd/goose@v3.24

.PHONY: deps
deps:
	go install $(MOCKERY_VERSION)
	docker pull $(GOLANGCI_LINT_IMAGE)
	go install $(GOOSE_VERSION)

.PHONY: mocks
mocks:
	$(MAKE) deps

	rm -rf ./internal/generated/mocks
	mockery

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
