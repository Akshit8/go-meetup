# go-meetup
A meet-up service designed in Go with GraphQL, CockroachDB and gqlgen

## Setting custom resolver on generated models with gqlgen
After the models are generated from provided SDL, wire them up as custom schema's by adding following to **gqlgen.yml** *models* key;
```
User:
  model: github.com/Akshit8/go-meetup/graph/model.User
    fields:
      meetups:
        resolver: true
Meetup:
  model: github.com/Akshit8/go-meetup/graph/model.Meetup
    fields:
      user:
        resolver: true
```

## Using Postgres for development
CockroachDB supports the PostgreSQL wire protocol, so available PostgreSQL client drivers and ORMs mostly work with CockroachDB. So for development we'll use postgres and we can later connect to cockroachDB cluster for production.
```bash
# setting postgres container
docker-compose up -d
docker exec -it postgresdb createdb --username=root --owner=root meetup

# setting migration with golang migrate
mkdir db/migration
migrate create -ext sql -dir db/migration -seq init

# migrate up
make migrationup
```

## Seeding PG db manually
```bash
docker exec -it postgresdb psql -U root -d meetup

INSERT INTO users (username, email) VALUES ('bob', 'bob@gmail.com');
INSERT INTO users (username, email) VALUES ('jon', 'jon@gmail.com');
INSERT INTO users (username, email) VALUES ('jane', 'jane@gmail.com');

INSERT INTO meetups (name, description, user_id) VALUES ('My first meetup', 'This is a description', 1);
INSERT INTO meetups (name, description, user_id) VALUES ('My second meetup', 'This is a description', 1);
```

## Optimising N + 1 queries with dataloaders
- Some GraphQL queries can make hundreds of database queries, often with mostly repeated data.
- Dataloader is a way to group up all of those concurrent requests, take out any duplicates, and store them in case they are needed later on in request. The dataloader is just that, a request-scoped batching and caching solution popularised by facebook.
- We’re going to use [dataloaden](https://github.com/vektah/dataloaden) to build our dataloaders.
```bash
go get github.com/vektah/dataloaden
mkdir dataloader
cd dataloader
go run github.com/vektah/dataloaden UserLoader string *github.com/Akshit8/go-meetup/graph/model.User
```

## Makefile specs
- **git** - git add - commit - push commands
- **start** - start the application without build
- **gen** - generated graphql-go code for graphql SDL
- **migrationUp** - migrate db to new migrations
- **migrationDown** - rollback db to previous stage

## References
[go-pg](https://medium.com/tunaiku-tech/go-pg-golang-postgre-orm-2618b75c0430) <br>

## Author
**Akshit Sadana <akshitsadana@gmail.com>**

- Github: [@Akshit8](https://github.com/Akshit8)
- LinkedIn: [@akshitsadana](https://www.linkedin.com/in/akshit-sadana-b051ab121/)

## License
Licensed under the MIT License