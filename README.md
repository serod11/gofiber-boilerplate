# Go REST API Boilerplate

This repository is a lightweight and structured Go (Golang) boilerplate codebase.
It follows best practices to ensure a clean and maintainable codebase.

The codebase provides the following features:

* Clean code
* Dependency injection
* Unit test
* Environment dependent application configuration management
* Database migration
* Live reloading during development

The codebase uses the following Go packages.

* Web framework: [gofiber](https://gofiber.io/)
* Database ORM: [gorm](https://gorm.io/)
* Dot env file reader: [godotenv](https://github.com/joho/godotenv)
* Redis client (optional): [go-redis](https://github.com/redis/go-redis)

## Project Layout

```
.
├── cmd                  main applications of the project
│   └── server           the API server application
├── pkg                  reusable packages
│   ├── config           config
│   ├── model            database entities
│   ├── repo             data access layer
│   ├── service          bussiness logic
│   └── utils            utility functions like initialize database connection 
└── test                 test files   
```

The top level directories `cmd`, `pkg` are commonly found in other popular Go projects,
also you can create `internal` directory to separate reusable packages into public and private packages.
[Standard Go Project Layout](https://github.com/golang-standards/project-layout).


## Getting Started

To run this project you have to have running postgres database and **Go 1.18 or above** installed locally. Then create `pkg/config/envs/.env` and set environment variables.

```shell
# download repository
git clone https://github.com/serod11/gofiber-boilerplate.git

cd gofiber-boilerplate

# run the REST API server
make run

# or run the API server with live reloading, which is useful during development
# requires air to be installed locally (https://github.com/cosmtrek/air)
make run-live

# build binary executable
make build
```

At this time, you have a REST API server running at `http://127.0.0.1:8080`. It provides the following endpoints:

* `GET /health`: a healthcheck service
* `GET /book`: returns all books from database book table
* `GET /book/:id`: returns a book by id

Try the URL `http://localhost:8080/healthcheck`, and you should see `{"status": "OK"}` displayed.


```shell
curl -X GET http://localhost:8080/health
```

## Authors
- [Satjan](https://github.com/satjan)
- [Ser-Od](https://github.com/serod11)
