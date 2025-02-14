swag:
	swag init --exclude docker,assets,pkg --md ./docs --parseInternal --parseDependency --parseDepth 2 -g ./cmd/app/main.go
# $ Update code s
wire:
	google-wire ./internal/di

docker-local: wire swag
	@docker compose --env-file ./.env.local -f ./local.docker-compose.yml -p iod up --build -d
	@$(MAKE) migrate-up-docker-local

### MIGRATIONS ###
migration-create:
	migrate create -ext sql -dir .\cmd\migrator\migrations -seq $(filter-out $@,$(MAKECMDGOALS))
migrate-up-docker-local:
	@docker compose --env-file ./.env.local -f ./local.docker-compose.yml -p iod run --rm migrate
	@docker compose --env-file ./.env.local -f ./local.docker-compose.yml -p iod run --rm migrate-mock
migrate-down-docker-local:
	@docker compose --env-file ./.env.local -f ./local.docker-compose.yml -p iod run --rm migrate-down-mock
	@docker compose --env-file ./.env.local -f ./local.docker-compose.yml -p iod run --rm migrate-down
