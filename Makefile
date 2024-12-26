.PHONY: postgres createdb dropdb migrateup sqlc test migratedown

DB_CONTAINER_NAME=postgres
DB_USER=root
DB_PASSWORD=root
DB_NAME=simplebank
DB_PORT=5432

postgres:
	@if docker ps -a --format '{{.Names}}' | grep -q "^${DB_CONTAINER_NAME}$$"; then \
		if docker ps --format '{{.Names}}' | grep -q "^${DB_CONTAINER_NAME}$$"; then \
			echo "Stopping container ${DB_CONTAINER_NAME}..."; \
			docker stop "${DB_CONTAINER_NAME}"; \
		fi; \
		echo "Removing container ${DB_CONTAINER_NAME}..."; \
		docker rm "${DB_CONTAINER_NAME}"; \
	else \
		echo "Container ${DB_CONTAINER_NAME} does not exist"; \
	fi;\
	docker run --name $(DB_CONTAINER_NAME) -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -h localhost -p $(DB_PORT):5432 -d postgres

createdb:
	docker exec -it $(DB_CONTAINER_NAME) createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

dropdb:
	docker exec -it $(DB_CONTAINER_NAME) dropdb --username=$(DB_USER) $(DB_NAME)

migrateup:
	migrate -path db/migration -database "postgresql://$(DB_USER):$(DB_PASSWORD)@127.0.0.1:$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

migratedown:
	migrate -path db/migration -database "postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down