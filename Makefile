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

docker-migrate-up:
	docker compose exec api go run cmd/migrations/main.go -direction up

docker-migrate-down:
	docker compose exec api go run cmd/migrations/main.go -direction down

docker-migrate-reset:
	docker compose exec api go run cmd/migrations/main.go -direction reset

docker-migrate-reset-up:
	docker compose exec api go run cmd/migrations/main.go -direction reset
	docker compose exec api go run cmd/migrations/main.go -direction up
