# Golang Basic Auth API

This project is a simple implementation of a basic authentication & authorization API using Go.

## Features

- User login and registration
- CRUD operations for users and todos
- Token-based authentication
- Role-based access control

## Requirements

- Go 1.23 or higher
- PostgreSQL

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/sherwin-77/golang-basic-auth-api.git
    ```
2. Navigate to the project directory:
    ```sh
    cd golang-basic-auth-api
    ```
3. Install dependencies:
    ```sh
    go mod tidy
    ```

## Configuration

Copy the `.env.example` file to `.env` and update the values as needed.

```sh
cp .env.example .env
```

## Running Migrations

Database migrations are handled using golang-migrate CLI. You might need to [install it](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) first

Run the migration using the following command:

```sh
make migrate
```

To rollback the migration, run:

```sh
make rollback step=<number_of_steps>
```

Additionally, you can create a new migration file using the following command:

```sh
make migration name=<migration_name>
```

## Usage

1. Run the application:
    ```sh
    make serve
    ```

2. By default, the application will run on port http://localhost:8080

## Seeding

To seed the database with a test user with roles, run the following command:

```sh
make seed
```

Default credentials for the test user are:

- Email: `user@example.com` | Password: `secret` (auth level 1)

- Email: `editor@example.com` | Password: `secret` (auth level 2)

- Email: `admin@example.com` | Password: `secret` (auth level 3)

## Endpoints
All endpoints are prefixed with `/api`

### User
All user endpoints are prefixed with `/user` (e.g. `/user/todos`)

Accessing user endpoints requires at least role with auth level 2

- `GET /todos` - Get all todos
- `POST /todos` - Create a new todo
- `GET /todos/:id` - Get a todo by ID
- `PUT /todos/:id` - Update a todo by ID
- `DELETE /todos/:id` - Delete a todo by ID

### Auth
All auth endpoints are prefixed with `/auth` (e.g. `/auth/login`)

- `POST /register` - Register a new user
- `POST /login` - Login 

### Admin
All admin endpoints are prefixed with `/admin` (e.g. `/admin/users`)

Accessing admin endpoints requires at least role with auth level 3

- `GET /users` - Get all users
- `POST /users` - Create a new user
- `GET /users/:id` - Get a user by ID
- `PUT /users/:id` - Update a user by ID
- `DELETE /users/:id` - Delete a user by ID
