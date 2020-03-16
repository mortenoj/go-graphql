# Go GraphQL Template

> GraphQL API template written in Golang for learning purposes

This is an example project showcasing how you can setup GraphQL in Golang using [Golang standard project structure](https://github.com/golang-standards/project-layout) and tools like [GQLGen](https://github.com/99designs/gqlgen), [Gin](https://github.com/gin-gonic/gin), [Realize](https://github.com/oxequa/realize), and [GORM](https://gorm.io/).

## Config

The program will look for the following environment variables:

```bash
# General configs
PRODUCTION  # {bool} {default: "true"} Program is in production mode or is running i a cluster etc...

# Server configs
GIN_MODE            # {string}      {default "debug"}       Gin mode either release or debug
GQL_SERVER_HOST     # {string}      {default "localhost"}   Server host
GQL_SERVER_PORT     # {string|int}  {default "8080"}        Server port

# GQLGen configs
GQL_SERVER_GRAPHQL_PATH                 # {string}  {default "/graphql"}    Endpoint path to GraphQL API
GQL_SERVER_GRAPHQL_PLAYGROUND_ENABLED   # {bool}    {default "true"}        Enable GraphQL playground interface
GQL_SERVER_GRAPHQL_PLAYGROUND_PATH      # {string}  {default "/"}           Endpoint path to GraphQL playground

# ORM config
GORM_DIALECT            # {string}                      Database dialect i.e postgres, mysql, etc...
GORM_CONNECTION_DSN     # {string}                      Database connection string i.e "postgres://user:password@dbhost/dbname?sslmode=disable"
GORM_SEED_DB            # {bool}    {default "true"}    Enable database seeding on run
GORM_LOGMODE            # {bool}    {default "true"}    Enable GORM log mode
GORM_AUTOMIGRATE        # {bool}    {default "true"}    Automatically run migrations on run
```

Set these variables in a `.env` file and they will be exported when running `make run-dev`. (See `scripts/run.sh`)

## Building

Simply run `make bin` to build the go binary to `bin/package`.

To build docker image run `make build`.

## Dependencies

Dependencies are handled by Go Modules and should be handled automatically when building. To comfirm run `make bin` to build.

If there is an issue run `go mod tidy` to clean up the mod file.

If your issue is regarding fetching private repos such as `go-utils` try to run `go get -insecure <repo>`, but it will persists until you fix your gitconfig and credentials.

## Running

If you have installed dependencies and are able to build, you can run the program by running `make run-dev`.
