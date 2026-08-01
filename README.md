# Go-Bank

A backend banking API built in Go with Fiber, PostgreSQL, SQLC, authentication, transactional transfers, tests, Docker, and CI/CD.

## Technical work handled

- Designed a banking data model in PostgreSQL (`accounts`, `entries`, `transfers`, `users`) with keys, indexes, constraints, and migrations.
- Built SQL-first data access with SQLC-generated queries and types.
- Implemented database access with `pgx`/`pgxpool`.
- Added transactional money transfer logic (`TransferTx`) with ordered balance updates to reduce deadlock risk.
- Built REST APIs with Fiber v3 for:
  - User registration and login
  - Account creation, account lookup, and paginated account listing
  - Account-to-account transfer
- Added auth using PASETO tokens and bearer-token middleware.
- Added secure password hashing with bcrypt.
- Added request validation with `go-playground/validator`.
- Added environment/config loading via Viper (`DB_SOURCE`, `SERVER_ADDRESS`, `TOKEN_SYMMETRIC_KEY`, `TOKEN_DURATION`).
- Added containerization with Docker and runtime migration execution on startup.
- Added local orchestration with Docker Compose.
- Added developer automation with Makefile targets for DB, migrations, SQLC generation, tests, and server run.
- Added automated tests and test bootstrap for DB-backed test runs.
- Added GitHub Actions workflows for:
  - CI test pipeline (Postgres service + migrations + `go test`)
  - Build and push Docker image to AWS ECR on main branch.

## Project structure

- `cmd/api/main.go` — application entrypoint
- `internals/` — server setup, routers, handlers, middleware
- `db/migration/` — SQL migrations
- `db/query/` — SQL source files for SQLC
- `db/sqlc/` — generated queries/models + transactional store
- `token/` — token interface, payload, PASETO maker
- `util/` — config, password utilities, random helpers
- `.github/workflows/` — CI/CD workflows

## API routes

### User
- `POST /users/` — create user
- `POST /users/login` — login and get access token

### Accounts (auth required)
- `POST /accounts/` — create account
- `GET /accounts/:id` — get one account
- `GET /accounts/?page_id=&page_size=` — list accounts

### Transfers (auth required)
- `POST /transfers/` — create transfer

Use header:

`Authorization: ******

## Run locally

### 1) Start postgres
Use Docker Compose:

```bash
docker compose up -d postgres
```

Or use Makefile helper commands (as defined):

```bash
make postgres
make createdb
```

### 2) Run migrations

```bash
make migrateup
```

### 3) Start API

```bash
make server
```

## Test

```bash
make test
```

## Docker run

```bash
docker compose up --build
```

The API container runs DB migrations from `start.sh` before starting the server.
