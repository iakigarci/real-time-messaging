.PHONY: migrate-create migrate-up migrate-down migrate-status install-goose help \
        migrate-down-to migrate-redo

include .env
export

DB_URL=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable
MIGRATIONS_DIR=migrations/
GOOSE_BIN=$(HOME)/go/bin/goose

# Install goose if not found
install-goose:
	@which goose >/dev/null 2>&1 || { \
		echo "Installing goose..."; \
		go install github.com/pressly/goose/v3/cmd/goose@latest; \
	}

# Create a new migration
migrate-create: install-goose
	@read -p "Enter migration name: " name; \
	if [ -z "$$name" ]; then echo "Migration name cannot be empty!"; exit 1; fi; \
	$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) create "$$name" sql

# Run all migrations
migrate-up: install-goose
	$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" up

# Roll back the last migration
migrate-down: install-goose
	$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" down

# Roll back to a specific migration version
migrate-down-to: install-goose
	@read -p "Enter target migration version: " version; \
	if [ -z "$$version" ]; then echo "Version cannot be empty!"; exit 1; fi; \
	$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" down-to $$version

# Re-run the last migration (down + up)
migrate-redo: install-goose
	$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" redo

# Reset all migrations
migrate-reset: install-goose
	$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" reset

# Show migration status
migrate-status: install-goose
	$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" status

# Help command to list available targets
help:
	@echo "Available commands:"
	@echo "  install-goose    Install goose migration tool"
	@echo "  migrate-create   Create a new migration file"
	@echo "  migrate-up       Apply all migrations"
	@echo "  migrate-down     Rollback the last migration"
	@echo "  migrate-down-to  Rollback to a specific version"
	@echo "  migrate-redo     Re-run the last migration"
	@echo "  migrate-status   Show migration status"
