.PHONY: postgres createdb dropdb migrateup sqlc test migratedown stop-database remove-network

DB_CONTAINER_NAME ?= postgres
DB_USER ?= root
DB_PASSWORD ?= root
DB_NAME ?= simplebank
DB_PORT ?= 5432
NETWORK_NAME ?= postgres
DB_HOST ?= postgres
start: postgres createdb migrateup
postgres: stop-database create-network
	docker run --name $(DB_CONTAINER_NAME) --network=$(NETWORK_NAME) -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -p $(DB_PORT):5432 -d postgres

stop: migratedown stop-database
# Target for creating the database
createdb:
	# # Initialize retry_count and max_retries
	# retry_count=0
	# max_retries=10  # Set the maximum number of retries

	# # Wait until PostgreSQL is ready to accept connections
	# until docker exec postgres pg_isready -U root -h postgres; do \
	# 	echo "Waiting for database to be ready... Attempt $((retry_count + 1)) of $max_retries"; \
	# 	sleep 2; \
	# 	retry_count=$((retry_count + 1)); \

	# 	# Stop after max_retries
	# 	if [ $retry_count -ge $max_retries ]; then \
	# 		echo "Reached maximum retry limit ($max_retries). Exiting."; \
	# 		exit 1; \
	# 	fi; \
	# done

	# Run the createdb command
	sleep 10
	docker exec -e PGPASSWORD=$(DB_PASSWORD) $(DB_CONTAINER_NAME) createdb -U $(DB_USER) -h $(DB_CONTAINER_NAME) -O $(DB_USER) $(DB_NAME)
dropdb:
	docker exec -e PGPASSWORD=$(DB_PASSWORD) $(DB_CONTAINER_NAME) dropdb -U $(DB_USER)  -h $(DB_CONTAINER_NAME) -O $(DB_USER) $(DB_NAME) 

migrateup:
	migrate -path db/migration -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_CONTAINER_NAME):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

migratedown:
	
	migrate -path db/migration -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_CONTAINER_NAME):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down -all

create-network:
	@if docker network ls --format '{{.Name}}' | grep -q "^$(NETWORK_NAME)$$"; then \
		echo "Network $(NETWORK_NAME) already exists."; \
	else \
		echo "Creating network $(NETWORK_NAME)..."; \
		docker network create $(NETWORK_NAME); \
	fi

stop-database:
	@if docker ps -a --format '{{.Names}}' | grep -q "^${DB_CONTAINER_NAME}$$"; then \
		if docker ps --format '{{.Names}}' | grep -q "^${DB_CONTAINER_NAME}$$"; then \
			echo "Stopping container ${DB_CONTAINER_NAME}..."; \
			docker stop "${DB_CONTAINER_NAME}"; \
		fi; \
		echo "Removing container ${DB_CONTAINER_NAME}..."; \
		docker rm "${DB_CONTAINER_NAME}"; \
	else \
		echo "Container ${DB_CONTAINER_NAME} does not exist"; \
	fi; 

# Remove Docker network
remove-network:
	@if docker network ls --format '{{.Name}}' | grep -q "^$(NETWORK_NAME)$$"; then \
		echo "Removing network $(NETWORK_NAME)..."; \
		docker network rm $(NETWORK_NAME); \
	else \
		echo "Network $(NETWORK_NAME) does not exist."; \
	fi