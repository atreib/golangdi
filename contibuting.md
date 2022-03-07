# .env file

Rename the `.env.copy` to `.env` and set the variables values according to your environment. They are pretty much mirrored with the `docker-compose.yaml` file.

Remember to run `source .env` to apply the variables into your terminal. This is going to be required for some commands.

# Infrastructure

We're using [docker-compose](https://docs.docker.com/compose/) to manage our infrastructure (as database and other services). After installing docker-compose in your machine, just run `docker-compose up -d`.

> If you want to read more about Docker and Docker-compose, check out my material at [Notion](https://cloudy-marsupial-788.notion.site/CI-CD-115134e2ddeb4c5bb9273912a235dd06).

## Relational Database

- DBMS: PostgreSQL
- Database name: `postgres` (default)
- Username: `postgres` (default)
- Password: `123456` (defined on `docker-compose.yaml`)
- Port: `5432` (default) mapped to `54321` on host
- A container's volume is stored `/.data/sqldb`
- We're using [DBeaver](https://dbeaver.io/) as an IDE to manage the database

# Folder structure

- `cmd`: Project's entrypoint
- `entities`: Entities/models
- `infra`: Lowest lever of abstraction
  - `libs`: Libraries non-related to the project goal, but required (jwt, encryption, etc)
  - `repositores`: Repositories of entities
- `services`: Services related to the project goal
- `database`: All database-related files
  - `migrations`: Migrations
  - `.data`: Database volume (Not commited)

# Migrations (PostgreSQL)

We're using [golang-migrate](https://github.com/golang-migrate/migrate) to manage the migrations of our PSQL database. You can install the CLI by following the [lib's docs](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate). They also provide an awesome documentation about how to use the CLI with an PostgreSQL database, which you can find [here](https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md).

All migrations are stored in the `/migrations` directory.
Every migration creates two files: `.up.sql` and `.down.sql`. They're the basic up/down migrations concept.

- Creating a new migration: `migrate create -ext sql -dir db/migrations -seq {migration_name}`
- Running migrations: `migrate -database ${POSTGRESQL_URL} -path database/migrations up`
- Rollback migrations: `migrate -database ${POSTGRESQL_URL} -path database/migrations down`

> Our migrations use variables from the .env file, so don't forget to setup your .env file and run `source .env` before using the migrations' commands
