MIGRATIONS_PATH=cmd/migrate/migrations
DB_URL=postgres://admin:adminpassword@db:5432/social?sslmode=disable

# biar "create_users" nggak dianggap target yang harus dibangun
%::
	@:

.PHONY: migration
migration: ## Usage: make migration <name>
	@name=$(firstword $(filter-out $@,$(MAKECMDGOALS))); \
	test -n "$$name" || (echo "Usage: make migration <name>"; exit 1); \
	echo "creating migration: $$name"; \
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $$name

.PHONY: migrate-up
migrate-up:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

.PHONY: migrate-down
migrate-down:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down 1

.PHONY: db-refresh
db-refresh: migrate-down migrate-up sqlc-generate