run:
	go run ./cmd/api

up:
	docker compose up -d

down:
	docker compose down

migrate-up:
	docker compose run --rm migrate \
		-path=/migrations \
		-database "postgres://postgres:123@host.docker.internal:5432/train_radar?sslmode=disable" \
		up

migrate-down:
	docker compose run --rm migrate \
		-path=/migrations \
		-database "postgres://postgres:123@host.docker.internal:5432/train_radar?sslmode=disable" \
		down 1