migrate-up:
	migrate -database "mysql://$(shell echo $$DB_USERNAME):$(shell echo $$DB_PASSWORD)@$(shell echo $$DB_HOST):$(shell echo $$DB_PORT)/$(shell echo $$DB_NAME)" -path database/migrations up

migrate-down:
	migrate -database "mysql://$(shell echo $$DB_USERNAME):$(shell echo $$DB_PASSWORD)@$(shell echo $$DB_HOST):$(shell echo $$DB_PORT)/$(shell echo $$DB_NAME)" -path database/migrations down

compose-build:
	docker-compose build

compose-up:
	docker-compose up

