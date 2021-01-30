package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Akshit8/go-meetup/db"
	"github.com/Akshit8/go-meetup/graph"
	"github.com/Akshit8/go-meetup/graph/generated"
	"github.com/Akshit8/go-meetup/graph/model"
	"github.com/go-pg/pg/v10"
)

const defaultPort = "8080"

var meetups = []*model.Meetup{
	{
		ID:          "1",
		Name:        "A meetup",
		Description: "A description",
		UserID:      "1",
	},
	{
		ID:          "2",
		Name:        "A second meetup",
		Description: "A description",
		UserID:      "2",
	},
}

var users = []*model.User{
	{
		ID:       "1",
		Username: "Bob",
		Email:    "bob@gmail.com",
	},
	{
		ID:       "2",
		Username: "Jon",
		Email:    "jon@gmail.com",
	},
}

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

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		MeetupStore: newMeetupRepo,
		UserStore:   newUserRepo,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
