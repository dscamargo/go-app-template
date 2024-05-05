# Template Golang w/ cli commands

### Stack

- [Fiber](https://docs.gofiber.io/)
- [Bun ORM](https://bun.uptrace.dev/)
- [Urfave CLI](https://cli.urfave.org/)

### Commands
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
