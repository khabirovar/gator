# gator

## Requirements
- PostgreSQL
- Goose
- Go

### Installation
Before utility using need to install PostgreSQL version 15 or higher and run migrations in schema.
Instead installing PostgreSQL you can run db with docker-compose command 
```
docker compose up -d
```
After db started we need install migration utility
```
go install github.com/pressly/goose/v3/cmd/goose@latest
```
Then utility installed we can run migrations, which one in `sql/schema`
```
cd sql/schema
goose postgres "postgres://wagslane:@localhost:5432/gator" up
```
Utility use config file `.gatorconfig.json` in the home directory.
In config file have to write dsn of database for utility
```
{
  "db_url":"postgres://postgres:postgres@127.0.0.1:5432/gator?sslmode=disable"
}
```
to install utility of you can use command
```
go install github.com/khabirovar/gator@latest

### Using 
TODO: write using
