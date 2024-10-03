migrate-up:
	migrate -database "mysql://blog_user:blog_user@tcp(localhost:3333)/blog" -path database/migrations up

migrate-down:
	migrate -database "mysql://blog_user:blog_user@tcp(localhost:3333)/blog" -path database/migrations down

compose-build:
	docker-compose build

compose-up:
	docker-compose up

