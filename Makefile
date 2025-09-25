MIGRATIONS_DIR=migrations/sql

# create migration file: name=create_table_name
migration:
	@echo "Creating migration: $(name)"
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)