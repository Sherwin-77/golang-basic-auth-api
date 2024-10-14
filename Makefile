include .env

serve:
	go run ./cmd/server

migrate:
	migrate -path db/migrations -database "postgresql://$(PSQL_USER):$(PSQL_PASSWORD)@$(PSQL_HOST):$(PSQL_PORT)/$(PSQL_DATABASE)?sslmode=disable" -verbose up $(step)

migrate-force:
	migrate -path db/migrations -database "postgresql://$(PSQL_USER):$(PSQL_PASSWORD)@$(PSQL_HOST):$(PSQL_PORT)/$(PSQL_DATABASE)?sslmode=disable" -verbose force $(version)

migrate-rollback:
	migrate -path db/migrations -database "postgresql://$(PSQL_USER):$(PSQL_PASSWORD)@$(PSQL_HOST):$(PSQL_PORT)/$(PSQL_DATABASE)?sslmode=disable" -verbose down $(step)

migration:
	migrate create -ext sql -dir db/migrations $(name)

seed-role:
	go run ./cmd/seeders/role

seed-user:
	go run ./cmd/seeders/user

seed:
	make seed-user
	make seed-role