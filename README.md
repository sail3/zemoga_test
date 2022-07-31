# zemoga_test Microservice

This repository contains a solutions to zemoga start test.

```bash
make init
make build
make up
```

And you are ready to Go!

`init` will populate the `.env` file needed for injecting environment variables. Make sure to have all the values as you need them.
`build` will create the development image to code inside of it.
`up` will run the API, exposing ports specified in the docker-compose file.

This lets the developer focus on the code, running it inside the container resembling production.

---

## Unit Testing

The unit testing was done using the dependency injection technique.

Same as the development, the unit testing is performed inside the docker container, to do so run the following:

```bash
make devshell
make t
```

`devshell` will run the development container and start a terminal inside of it.
`t` will run the unit testing and provide the coverage level for each package.

---

## Database

We use **PostgreSQL** as the database for the service. The configuration should work out of the box, but if running the service for the first time you will need to run the migrations.

The **docker-compose** already has a DB that will be setup when you run `make up`.

### Running Database Migrations

As everything else, migrations must be run inside the container. This enables consistency and allows for the configuration to work as expected.

```bash
make devshell
make migrate-up
```

`devshell` will run the development container and start a terminal inside it.
`migrate-up` will run the database migrations located in the *migrations* folder.

### Adding a new Migration

We are using [migrate](https://github.com/golang-migrate/migrate) to handle migrations.
So in order to create a new migration properly, we will need to execute the following commands:

```bash
make devshell
migrate create -ext sql -dir migrations -seq name_of_migration
```

This will create the `up` and `down` migrations for the `name_of_migration`, then you can write the necessary commands in those.

---
