package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Akshit8/go-meetup/domain"

	auth "github.com/Akshit8/go-meetup/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Akshit8/go-meetup/db"
	"github.com/Akshit8/go-meetup/graph/dataloader"
	"github.com/Akshit8/go-meetup/graph/generated"
	"github.com/Akshit8/go-meetup/graph/resolver"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-pg/pg/v10"
	"github.com/rs/cors"
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

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

	router.Use(auth.AuthMiddleware(newUserRepo))

	d := domain.NewDomain(newUserRepo, newMeetupRepo)

	c := generated.Config{Resolvers: &resolver.Resolver{Domain: d}}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", dataloader.Middleware(dbConn, srv))

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
