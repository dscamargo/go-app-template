# Template Golang and Bun ORM w/ cli commands

### Stack

- [Air - Live reload for Go apps](https://github.com/cosmtrek/air)
- [Fiber](https://docs.gofiber.io/)
- [Bun ORM](https://bun.uptrace.dev/)
- [Urfave CLI](https://cli.urfave.org/)

### CLI Commands

- go run cmd/cli/main.go [command]

| Command          | Description                                                  |
|------------------|--------------------------------------------------------------|
| app run          | Run fiber web server                                         |
| db init          | Create orm default databases (locks, migrations table, etc.) |
| db migrate       | Run database migrations (migrations folder)                  |
| db rollback      | Rollback migrations                                          |
| db create_go     | Create migration in golang format                            |
| db create_sql    | Create migration in SQL format                               |
| db create_tx_sql | Create transactional migration in SQL format                 |
| db status        | Show database migrations status                              |

### Makefile commands

- make [command]

| Command       | Description                                              |
|---------------|----------------------------------------------------------|
| build         | Golang build                                             |
| run-db        | Run docker compose database service                      |
| compose-build | Run docker compose build                                 |
| exec-bash     | Run docker compose exec in app service                   |
| run-dev       | Run application in dev mode with live reload|

### How to use

- Create a repository choosing the template
- Rename go module
- Run `make run-db` to run database service
- Run `make compose-build` to build the app
- Run `make run-dev` to start the app and attach in container
- Enjoy !