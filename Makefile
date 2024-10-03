start:
	go run cmd/SAProject/main.go

gen:
	sqlc generate
	wire ./internal/...

docker-gen:
	docker compose exec api make gen

gen-sqlc:
	sqlc generate

docker-gen-sqlc:
	docker compose exec api make gen-sqlc

gen-wire:
	wire ./internal/...

docker-gen-wire:
	docker compose exec api make gen-wire

migrate-schema:
	go run cmd/migration/main.go --migrate:schema

migrate-up:
	go run cmd/migration/main.go --migrate:up

migrate-down:
	go run cmd/migration/main.go --migrate:down --step=$(step)

migrate-reset:
	go run cmd/migration/main.go --migrate:reset

migrate-make:
	go run cmd/migration/main.go --migrate:make --name=$(name)

docker-migrate-up:
	docker compose exec api goose -dir ./cmd/migrations postgres "host=postgres user=myuser dbname=mydatabase password=mypassword port=5432 sslmode=disable" up
