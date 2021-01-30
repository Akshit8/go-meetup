git:
	git add .
	git commit -m "$(msg)"
	git push origin master

gen:
	go run github.com/99designs/gqlgen generate

start:
	go run server.go

.PHONY: git gen start