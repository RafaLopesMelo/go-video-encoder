include internal/infra/config/env/.env

MIGRATION_NAME ?=
DATABASE_URL = "postgresql://$(POSTGRES_USERNAME):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=disable"

create_migration:
	@if [ -z "$(MIGRATION_NAME)" ]; then \
		echo "You must pass the migration name through the variable 'MIGRATION_NAME'."; \
		exit 1; \
	else \
		echo "Creating migration: $(MIGRATION_NAME)"; \
		./scripts/migrate create -ext=sql -dir=internal/database/migration -seq $(MIGRATION_NAME); \
	fi

migrate_up:
	@echo "Running migrations..."
	./script/migrate -path=internal/database/migration -database $(DATABASE_URL) -verbose up 
	@echo "Migrations ran successfully!"

migrate_down:
	@echo "Running rollback..."
	./script/migrate -path=internal/database/migration -database $(DATABASE_URL) -verbose down 1
	@echo "Rollback executed successfully!"

