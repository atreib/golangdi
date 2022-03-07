# API with Dependency Injection Example

A brief example of dependency injection in an API with Golang.
Run `go run cmd/server.go` to execute the example.

We're going to implement some concepts, more related to API development, using Golang.
Get ready to see some implementation of the SOLID principals, with DI included, and GraphQL, RESTful, and gRPC endpoints, in order to approach the main communication protocols using Go.

At last but not at least, we're also implementing some different real-world services, as connecting with relational database (PostgreSQL), with non-relational databases (MongoDB) and some message service (probably RabbitMQ, I didn't thought about that yet)...

I think that you've already got the idea. I want to to do a few of everything using Go.

Keep reading on `contibuting.md`.

# Get started

We're using [docker-compose](https://docs.docker.com/compose/) to manage our infrastructure (as database and other services). After installing docker-compose in your machine, just run `docker-compose up -d` to set up everything.

Another required step is to set your `.env` file. So, create a copy of the `.env.copy` file and rename it to `.env`. It's going to be pretty much mirrored with your `docker-compose.yaml`. If you haven't changed anything, then you won't need to change anything in your `.env` file. You can read more about it in the `contributing.md` file.

With docker-compose correcly installed and the `.env` file correctly set, let's load the `.env` file variables into your terminal, by running `source .env`. This is required in order to run our migrations. So, as you can already imagine, let's run our migrations: `migrate -database ${POSTGRESQL_URL} -path database/migrations up`.

Now you're ready to run our API (`go run cmd/server.go`).
Feel free to read the `contributing.md` file and improve this project.

# Useful commands

- Create/start the infrastructure containers: `docker-compose up -d`
- Load the .env file into shell: `source .env`
- Start the application: `go run cmd/server.go`
- Create a new migration: `migrate create -ext sql -dir db/migrations -seq {migration_name}`
- Running migrations: `migrate -database ${POSTGRESQL_URL} -path database/migrations up`
- Rollback migrations: `migrate -database ${POSTGRESQL_URL} -path database/migrations down`
- Close the infrastructure containers: `docker-compose down`
- Run the unit tests: `go test -v ./...`
