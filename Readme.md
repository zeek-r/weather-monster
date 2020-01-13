# Weater Monster API

---

# Pre-requisites:

```
Golang 1.12 and above
Postgresql

```

## Libraries used:

    - Gorilla mux as http router
    - gorm as postgres ORM
    - Logrus as logger

## Project structure

The project structure tries to follow clean architecture

- api for all the transport layer related functions.
  - e.g. http handlers, routes, middlewares
- app for storing the business logic of the application
  - e.g. various services
- cmd for storing command line functions for starting app/api
- config for storing all the configurations files.
- db for storing db access functions.
- model for storing db models.
- pkg for shared functions that can be reused in other repos.
  - e.g. logger

## Instructions for setting-up project

    - Move .env.example to .env and change the configurations accordingly.
    - Install project dependencies

>     	go mod download

## How-to

---

### Run Project
- Dev mode
    > ./dev.sh
- Build
  > go build -o weather-monster .
- Run
  > ./weather-monster start
- Build and Run
  > go run ~/go/src/github.com/zeek-r/weather-monster start
- Test
  > go test ./... -coverprofile cover.out
- Test coverage analysis
  > go tool cover -html=cover.out
