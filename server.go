package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Akshit8/go-meetup/db"
	"github.com/Akshit8/go-meetup/graph/dataloader"
	"github.com/Akshit8/go-meetup/graph/generated"
	"github.com/Akshit8/go-meetup/graph/resolver"
	"github.com/go-pg/pg/v10"
)

const defaultPort = "8080"

func main() {
	dbSource := os.Getenv("DB_SOURCE")
	opt, err := pg.ParseURL(dbSource)
	if err != nil {
		log.Fatal("unable to parse db source: ", err)
	}

	dbConn := db.NewDBConnection(opt)
	defer dbConn.Close()
	dbConn.AddQueryHook(db.Logger{})

	newMeetupRepo := db.NewMeetupRepo(dbConn)
	newUserRepo := db.NewUserRepo(dbConn)

	c := generated.Config{Resolvers: &resolver.Resolver{
		MeetupStore: newMeetupRepo,
		UserStore:   newUserRepo,
	}}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", dataloader.Middleware(dbConn, srv))

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
