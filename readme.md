# Dependency Injection Example

A brief example of dependency injection with Golang.
Run `go run cmd/server.go` to execute the example.

## Folder structure

- `cmd`: Project's entrypoint
- `entities`: Entities/models
- `infra`: Lowest lever of abstraction
  - `libs`: Libraries non-related to the project goal, but required (jwt, encryption, etc)
  - `repositores`: Repositories of entities
- `services`: Services related to the project goal