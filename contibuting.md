# Get started

You can set you environment through [docker-compose](https://docs.docker.com/compose/).
After installing it in your machine, just run `docker-compose up`.

> If you want to read more about Docker and Docker-compose, check out my material at [Notion](https://cloudy-marsupial-788.notion.site/CI-CD-115134e2ddeb4c5bb9273912a235dd06).

# Infrastructure

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
- `.data`: Local store (Not commited)
