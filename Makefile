git:
	git add .
	git commit -m "$(msg)"
	git push origin master

gen:
	go run github.com/99designs/gqlgen generate

start:
	go run server.go

migrationup:
	migrate -path db/migration -database $(DB_SOURCE) -verbose up

migrationdown:
	migrate -path db/migration -database $(DB_SOURCE) -verbose down

.PHONY: git gen start migrationup migrationdown