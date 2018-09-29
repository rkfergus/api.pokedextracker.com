# api.pokedextracker.com

[![CircleCI](https://circleci.com/gh/pokedextracker/api.pokedextracker.com.svg?style=shield)](https://circleci.com/gh/pokedextracker/api.pokedextracker.com)

The API for [pokedextracker.com](https://pokedextracker.com).

## Install

This project is meant to be run with at least Go v1.10 (due to new features
added to coverage tracking). If you use
[`goenv`](https://github.com/syndbg/goenv), it should pick up the `.go-version`
file in this repo and use the appropriate version. It also currently uses
[`dep`](https://github.com/golang/dep) to manage dependencies. If you're running
the project for the first time, you can run the commands below to get everything
set up correctly.

This project uses PostgreSQL 9.x as its database. Assuming you already have it
installed (either through [`brew`](http://brew.sh/) on OS X or `apt-get` on
Ubuntu), you can run `make setup` to setup the role and database.

```sh
$ goenv install 1.10.4
$ mkdir -p $GOPATH/src/github.com/pokedextracker
$ git clone git@github.com:pokedextracker/api.pokedextracker.com.git $GOPATH/src/github.com/pokedextracker/api.pokedextracker.com
$ cd $GOPATH/src/github.com/pokedextracker/api.pokedextracker.com
$ make setup
$ make install
```

### Database

Now that the database and role are setup, you can just run the following
commands to run the migrations.

```sh
$ make migrate
```

### Secrets

There are some secrets needed to run this repo locally, such as the Stripe API
Key. Since no secrets are being checked in, you should copy `.env.example` to
`.env` and populate it with all of the secrets listed there.

## Data

This repo doesn't include a way to completely load up the DB with all of the
actual Pokemon data. That's only been loaded into the staging and production
databases. For testing purposes and to make sure everything is functioning as
expected, having that data isn't entirely necessary. You should be relying on
tests and factories instead of the database state.

## Development

To run the server locally during active development, you can just run the
following command.

```sh
$ make serve
```

This will start the server listening on port 8648. So to interact with it, you
can use `curl`.

```sh
$ curl -s http://localhost:8648/pokemon
```

## Tests

To run the tests and collect coverage information, you can just run the
following command.

```sh
$ make test
```

It will output the results of the test, and a coverage summary for each package.
To see an HTML breakdown of coverage to see what you missed, you should run the
following command.

```sh
$ make html
```

To see the cumulative coverage total and find out if it meets the project
requirements, you can run the following command.

```sh
$ make enforce
```

## Docker

Every merge into the `master` branch on GitHub triggers a new build for a Docker
image. That image will overwrite the `latest` tag, and there will be an explicit
tag with the first 7 characters of the commit hash. The server will be listening
on port 8648 so if you run a container locally, make sure that traffic is
forwarded to that port. For example:

 ```sh
$ docker run --rm --publish 8648:8648 --name pokedextracker-api pokedextracker/api.pokedextracker.com:latest
```
