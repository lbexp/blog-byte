# Blog-byte

API for blog appication.
This project is only for assignment purpose.

## Requirements

- go v1.21.0
- docker
- golang-migrate/migrate (for db migrations)

## How to run locally

- Run docker compose build

  ```bash
  Make compose-build
  ```

- Run docker compose up

  ```bash
  Make compose-up
  ```

- Run migration for the database

  ````bash
  migrate -database "mysql://blog_user:blog_user@tcp(localhost:3333)/blog" -path database/migrations up```
  ````
