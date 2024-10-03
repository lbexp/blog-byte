# ✍️ Blog-byte

API for blog application.

This project is only for assignment purpose.

## 💡 Endpoints

### User

- `POST api/v1/register` - Register a new user.
- `POST api/v1/login` - Login and receive a token for authentication.

### Posts

- `POST api/v1/posts` - Create a new blog post.
- `GET api/v1/posts/{id}` - Get blog post details by ID.
- `GET api/v1/posts` - List all blog posts.
- `PUT api/v1/posts/{id}` - Update a blog post.
- `DELETE api/v1/posts/{id}` - Delete a blog post.

### Comments

- `POST api/v1/posts/{id}/comments` - Add a comment to a blog post.
- `GET api/v1/posts/{id}/comments` - List all comments for a blog post.

## 📕 Database Schema Design

<img title="Database Schema Design" alt="Database Schema Design" src="/db-schema.png">

Detail can be seen in [here](https://dbdiagram.io/d/Blog-byte-DB-Schema-66feecf5fb079c7ebd442bff) or in the `database/migrations` files.

## 📖 Requirements

- go v1.21.0
- docker & docker-compose
- golang-migrate/migrate (for db migrations)

## ⚙️ How to run locally

- Run docker compose build

  ```bash
  Make compose-build
  ```

- Run docker compose up

  ```bash
  Make compose-up
  ```

- Run migration for the database (one time only)

  ```bash
  Make migrate-up
  ```

## 📋 Other commands

- Rollback migration

  ```bash
  Make migrate-down
  ```

- Clean up service build caches

  ```bash
  ./clean.sh
  ```
